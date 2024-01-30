/*
Copyright © 2024 lflxp <382023823@qq.com>
*/
package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	. "github.com/lflxp/smkubectl/completion"
	"github.com/spf13/cobra"
)

func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// 解析命令
// 只负责执行 不负责命令补全
func parseCmd(in string) {
	var err error
	tmp := strings.Split(in, " ")
	result := []string{}
	// 记录命令最后一个字符是否是空格
	isLastWorkSpace := false
	// 清除空格和空字符串
	for index, v := range tmp {
		if v != "" && v != " " {
			result = append(result, v)
		}
		if index == len(tmp)-1 {
			if v == "" || v == " " {
				isLastWorkSpace = true
				// fmt.Printf("0 islaster %d %d works |%s| %v\n", index, len(tmp)-1, v, isLastWorkSpace)
			}
		}
	}

	if len(result) == 0 {
		return
	}

	if value, ok := Completes[result[0]]; ok {
		// 判断一级命令
		if len(result) == 1 {
			if value.IsShell {
				// fmt.Println("one", value.Shell)
				execCmd(value.Shell)
			} else {
				// fmt.Println("two", in, result)
				fmt.Println(strings.Join(value.Cmd, "\n"))
				fmt.Println(strings.Join(value.Args, "\n"))
			}
		} else {
			// 判断一级后面的最后一个命令
			// fmt.Println("three")
			if value_daughter, ok := value.Daughter[result[len(result)-1]]; ok {
				// 仅针对第一级命令有效
				rs := []string{}
				rs, err := execCompletion(rs, result, "", &value_daughter, -1, false, true)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Printf(strings.Join(rs, "\n"))
			} else {
				// fmt.Println("six")
				// 根据命令长度智能补全命令
				if len(result) == 2 {
					t_one := []string{}
					for index, k := range value.Cmd {
						if index == 0 {
							found, err := regexp.MatchString("(COMMAND|NAME|NAMESPACE).*", k)
							if err != nil {
								fmt.Println(err.Error())
								return
							}

							if !found {
								fmt.Println("COMMAND")
							} else {
								fmt.Println(k)
							}
						}
						if strings.HasPrefix(k, result[1]) {
							t_one = append(t_one, strings.Replace(k, result[1], "", 1))
						}
					}
					fmt.Printf(strings.Join(t_one, "\n"))
				} else if len(result) == 3 && !isLastWorkSpace {
					// 补全命令
					t_two := []string{}
					if value_3, ok := value.Daughter["get"]; ok {
						t_two, err = execCompletion(t_two, result, "", &value_3, 1, true, false)
						if err != nil {
							fmt.Println(err.Error())
							return
						}
					}
					fmt.Printf(strings.Join(t_two, "\n"))
				} else {
					// 补全数据
					t_two := []string{}
					// 优先匹配value.Daughter里面的值
					// 补全上一个命令的结果，查询并替换已有数据
					// 没有发现命令则获取上一步的命令 获取资源后进行资源匹配
					if len(t_two) == 0 && isLastWorkSpace {
						// 补全实时数据结果
						// 补全数据 有空格
						// 如果获取最后一个参数无数据 则执行该命令
						cmd := fmt.Sprintf("kubectl get %s -A", result[len(result)-1])
						t_two, err = execCompletion(t_two, result, cmd, nil, 1, true, false)
						if err != nil {
							fmt.Println(err.Error())
							return
						}
					} else if len(t_two) == 0 && !isLastWorkSpace {
						// 补全命令
						if target, ok := value.Daughter["get"]; ok {
							t_two, err = execCompletion(t_two, result, "", &target, 1, true, false)
							if err != nil {
								fmt.Println(err.Error())
								return
							}
						} else {
							fmt.Printf("key get not exist\n")
							return
						}

						// 补全可能缺失的数据
						if value_maybe, ok := value.Daughter[result[len(result)-2]]; ok {
							t_two, err = execCompletion(t_two, result, "", &value_maybe, 1, true, false)
							if err != nil {
								fmt.Println(err.Error())
								return
							}
						}
						// fmt.Println("15")
						// 补全命令无效 获取上级命令的结果 并补全数据prefix数据
						if len(t_two) == 0 {
							if value_daughter2, ok := value.Daughter[result[len(result)-2]]; ok {
								t_two, err = execCompletion(t_two, result, "", &value_daughter2, 2, true, false)
								if err != nil {
									fmt.Println(err.Error())
									return
								}
							}
						}
					} else {
						// fmt.Println("12")
						fmt.Printf("t_two %v %b\n", t_two, isLastWorkSpace)
					}
					fmt.Printf(strings.Join(t_two, "\n"))
				}
			}
		}
	} else {
		// fmt.Println("7")
		// execCmd(in)
	}
}

// 抽离namespace和name
func absoftNS(in []string) (string, string, error) {
	var namespace, name string
	var err error

	for index, value := range in {
		if value == "-n" {
			if len(in) > index+1 {
				namespace = in[index+1]
			} else {
				err = fmt.Errorf("namespace is empty")
				return namespace, name, err
			}

			if len(in) > index+2 {
				name = in[index+2]
				break
			} else {
				err = fmt.Errorf("name is empty")
				return namespace, name, err
			}
		}
	}
	return namespace, name, nil
}

// 执行Completion
// cmd不为空时 ==> isLastWorkSpace == true
// cmd为空时 ==> isLastWorkSpace == false
// keepHeader 是否保留第一行数据
// in 当前输入的数据
// needRawCmd 是否需要原始命令
func execCompletion(result, in []string, cmd string, daughter *Completion, target int, keepHeader, first bool) ([]string, error) {
	if daughter != nil && cmd == "" {
		if daughter.IsCondition {
			var (
				found bool
				err   error
			)
			// 多条件过滤
			for _, v := range daughter.Condition {
				found, err = regexp.MatchString(v, strings.Join(in, " "))
				if err != nil {
					return result, err
				}
			}
			if found {
				var shell string
				// 硬编码
				// 专门处理pod类查询
				switch daughter.Level {
				case "show containers":
					// 获取pod的namespace和name
					ns, name, err := absoftNS(in)
					if err != nil {
						return result, err
					}
					shell = fmt.Sprintf("kubectl get pod -n %s %s -o jsonpath='{.spec.containers[*].name}'|tr ' ' '\\n'", ns, name)
					fmt.Println("CONTAINERS")
				default:
					shell = daughter.Shell
				}

				execCmd(shell)
			}
		} else {
			if daughter.IsShell {
				// 实时补全数据结果
				if first {
					execCmd(daughter.Shell)
				} else {
					rs_string, err := execCmdString(daughter.Shell)
					if err != nil {
						return result, err
					}
					// 补全 -n中提供的数据后缀
					for index, v := range strings.Split(rs_string, "\n") {
						if keepHeader && index == 0 {
							result = append(result, v)
						} else {
							if strings.HasPrefix(v, in[len(in)-1]) {
								result = append(result, strings.Replace(v, in[len(in)-1], "", 1))
							}
						}
					}
				}
			} else {
				// 补全cmd命令后缀
				if first {
					if daughter.Cmd != nil {
						fmt.Println(strings.Join(daughter.Cmd, "\n"))
					}
					if daughter.Args != nil {
						fmt.Println(strings.Join(daughter.Args, "\n"))
					}
				} else {
					for _, k := range daughter.Cmd {
						if strings.HasPrefix(k, strings.TrimSpace(in[len(in)-target])) {
							t_rs := strings.Replace(k, strings.TrimSpace(in[len(in)-target]), "", 1)
							if t_rs != "" && t_rs != " " {
								result = append(result, t_rs)
							}
						}
					}
					for _, k := range daughter.Args {
						if strings.HasPrefix(k, strings.TrimSpace(in[len(in)-target])) {
							t_rs := strings.Replace(k, strings.TrimSpace(in[len(in)-target]), "", 1)
							if t_rs != "" && t_rs != " " {
								result = append(result, t_rs)
							}
						}
					}
				}
			}
		}
	} else if daughter == nil && cmd != "" {
		rs_string, err := execCmdString(cmd)
		if err != nil {
			// fmt.Println("9,1 ", err.Error())
			// 执行命令错误时执行整个命令
			if strings.Contains(err.Error(), "exit") {
				// fmt.Println("12 ", in, result)
				CMD := strings.Join(in, " ")
				if !strings.Contains(CMD, "edit") && in[target] != "edit" {
					execCmd(CMD)
				} else {
					execCmd(strings.Replace(CMD, "edit", "get", 1))
				}
			} else {
				return result, err
			}
		} else {
			// fmt.Println("9.2", t_two)
			// 补全 -n中提供的数据后缀
			for _, v := range strings.Split(rs_string, "\n") {
				if strings.TrimSpace(v) != "" {
					result = append(result, v)
				}
			}
		}
	} else {
		return result, fmt.Errorf("unsupport type cmd %s", cmd)
	}
	return result, nil
}

func execCmdString(in string) (string, error) {
	cmd := exec.Command("sh", "-c", in)
	cmd.Stdin = os.Stdin
	var out bytes.Buffer
	cmd.Stdout = &out // 标准输出
	var errout bytes.Buffer
	// cmd.Stderr = os.Stderr // 标准错误
	cmd.Stderr = &errout // 标准错误
	err := cmd.Run()
	if err != nil {
		// if &errout != nil {
		// 	log.Fatalf("cmd.Run() failed with %s err: %s errout: %s\n", in, err, errout.String())
		// } else {
		// 	log.Fatalf("cmd.Run() failed with %s err: %s\n", in, err)
		// }

		return "", err
	}
	return out.String(), nil
}

func execCmd(in string) error {
	cmd := exec.Command("sh", "-c", in)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout // 标准输出
	// var errout bytes.Buffer
	cmd.Stderr = os.Stderr // 标准错误
	// cmd.Stderr = &errout // 标准错误
	err := cmd.Run()
	if err != nil {
		// if &errout != nil {
		// 	log.Fatalf("cmd.Run() failed with %s err: %s errout: %s\n", in, err, errout.String())
		// } else {
		// 	log.Fatalf("cmd.Run() failed with %s err: %s\n", in, err)
		// }

		return err
	}
	return nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "demo",
	Short: "A brief descriptsion of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// parseCmd(args[1:])
	// 	fmt.Println(strings.Join(args, " "))
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// 保证没有参数或者参数只有一个且为completion的时候执行cobra
	// 其余都走parseCmd
	if len(os.Args) >= 2 {
		if os.Args[1] == "completion" {
			err := rootCmd.Execute()
			if err != nil {
				os.Exit(1)
			}
		} else {
			parseCmd(strings.Join(os.Args[1:], " "))
		}
	} else {
		parseCmd(strings.Join(os.Args[1:], " "))
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.demo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
