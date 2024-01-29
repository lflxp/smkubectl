/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

const FZF_OPTION = "FZF_DEFAULT_OPTS=\"--height ${FZF_TMUX_HEIGHT:-50%} --min-height 15 --reverse $FZF_DEFAULT_OPTS --preview 'echo {}' --preview-window down:3:wrap $FZF_COMPLETION_OPTS\" fzf -m"

type Completion struct {
	Level    string
	Cmd      []string // command命令
	Args     []string // 参数
	IsShell  bool     // 是否执行kubectl获取还是直接cmd提示
	Shell    string   // 获取提示的命令
	Daughter map[string]Completion
}

var Completes map[string]Completion = map[string]Completion{
	"k": Completion{
		Level: "cmd",
		Cmd: []string{
			"explain",
			"get",
			"edit",
			"delete",

			"rollout",
			"scale",
			"autoscale",

			"certificate",
			"cluster-info",
			"top",
			"cordon",
			"uncordon",
			"drain",
			"taint",

			"describe",
			"logs",
			"attach",
			"exec",
			"port-forward",
			"proxy",
			"cp",
			"auth:",
			"debug",

			"diff",
			"apply",
			"patch",
			"replace",
			"wait",
			"kustomize",

			"label",
			"annotate",
			"completion",

			"api-resources",
			"api-versions",
			"config",
			"plugin",
			"version",
		},
		Args:    []string{},
		IsShell: false,
		Daughter: map[string]Completion{
			"get": Completion{
				Level:   "get",
				IsShell: true,
				Shell:   "kubectl api-resources|sed 1d",
				Cmd: []string{"bindings", "componentstatuses", "configmaps", "endpoints", "events", "limitranges", "namespaces", "nodes", "persistentvolumeclaims", "persistentvolumes", "pods", "podtemplates", "replicationcontrollers", "resourcequotas", "secrets", "serviceaccounts", "services", "challenges", "orders", "mutatingwebhookconfigurations", "validatingwebhookconfigurations", "customresourcedefinitions", "apiservices", "controllerrevisions", "daemonsets", "deployments", "replicasets", "statefulsets", "applications", "applicationsets", "appprojects", "argocdextensions", "tokenreviews", "localsubjectaccessreviews", "selfsubjectaccessreviews", "selfsubjectrulesreviews", "subjectaccessreviews", "horizontalpodautoscalers", "cronjobs", "jobs", "certificaterequests", "certificates", "clusterissuers", "issuers", "certificatesigningrequests", "devworkspaceoperatorconfigs", "devworkspaceroutings", "leases", "bgpconfigurations", "bgppeers", "blockaffinities", "caliconodestatuses", "clusterinformations", "felixconfigurations", "globalnetworkpolicies", "globalnetworksets", "hostendpoints", "ipamblocks", "ipamconfigs", "ipamhandles", "ippools", "ipreservations", "kubecontrollersconfigurations", "networkpolicies", "networksets", "endpointslices", "events", "flowschemas", "prioritylevelconfigurations", "departments", "licenses", "products", "projects", "roletemplatebindings", "roletemplates", "users", "workzooms", "clusterpolicies", "clusterreportchangerequests", "generaterequests", "policies", "reportchangerequests", "updaterequests", "nodes", "pods", "alertmanagerconfigs", "alertmanagers", "podmonitors", "probes", "prometheuses", "prometheusrules", "servicemonitors", "thanosrulers", "ingressclasses", "ingresses", "networkpolicies", "runtimeclasses", "checlusters", "poddisruptionbudgets", "clusterrolebindings", "clusterroles", "rolebindings", "roles", "priorityclasses", "csidrivers", "csinodes", "csistoragecapacities", "storageclasses", "volumeattachments", "clusterpolicyreports", "policyreports", "devworkspaces", "devworkspacetemplates",
					"cs", "cm", "ep", "ev", "limits", "ns", "no", "pvc", "pv", "po", "rc", "quota", "sa", "svc", "crd", "crds", "ds", "deploy", "rs", "sts", "app", "apps", "appset", "appsets", "appproj", "appprojs", "hpa", "cj", "cr", "crs", "cert", "certs", "csr", "dwoc", "dwr", "ev", "dep", "cpol", "crcr", "gr", "pol", "rcr", "ur", "amcfg", "am", "pmon", "prb", "prom", "promrule", "smon", "ruler", "ing", "netpol", "pdb", "pc", "sc", "cpolr", "polr", "dw", "dwt",
				},
			},
			"edit": Completion{
				Level:   "edit",
				IsShell: true,
				Cmd: []string{"bindings", "componentstatuses", "configmaps", "endpoints", "events", "limitranges", "namespaces", "nodes", "persistentvolumeclaims", "persistentvolumes", "pods", "podtemplates", "replicationcontrollers", "resourcequotas", "secrets", "serviceaccounts", "services", "challenges", "orders", "mutatingwebhookconfigurations", "validatingwebhookconfigurations", "customresourcedefinitions", "apiservices", "controllerrevisions", "daemonsets", "deployments", "replicasets", "statefulsets", "applications", "applicationsets", "appprojects", "argocdextensions", "tokenreviews", "localsubjectaccessreviews", "selfsubjectaccessreviews", "selfsubjectrulesreviews", "subjectaccessreviews", "horizontalpodautoscalers", "cronjobs", "jobs", "certificaterequests", "certificates", "clusterissuers", "issuers", "certificatesigningrequests", "devworkspaceoperatorconfigs", "devworkspaceroutings", "leases", "bgpconfigurations", "bgppeers", "blockaffinities", "caliconodestatuses", "clusterinformations", "felixconfigurations", "globalnetworkpolicies", "globalnetworksets", "hostendpoints", "ipamblocks", "ipamconfigs", "ipamhandles", "ippools", "ipreservations", "kubecontrollersconfigurations", "networkpolicies", "networksets", "endpointslices", "events", "flowschemas", "prioritylevelconfigurations", "departments", "licenses", "products", "projects", "roletemplatebindings", "roletemplates", "users", "workzooms", "clusterpolicies", "clusterreportchangerequests", "generaterequests", "policies", "reportchangerequests", "updaterequests", "nodes", "pods", "alertmanagerconfigs", "alertmanagers", "podmonitors", "probes", "prometheuses", "prometheusrules", "servicemonitors", "thanosrulers", "ingressclasses", "ingresses", "networkpolicies", "runtimeclasses", "checlusters", "poddisruptionbudgets", "clusterrolebindings", "clusterroles", "rolebindings", "roles", "priorityclasses", "csidrivers", "csinodes", "csistoragecapacities", "storageclasses", "volumeattachments", "clusterpolicyreports", "policyreports", "devworkspaces", "devworkspacetemplates",
					"cs", "cm", "ep", "ev", "limits", "ns", "no", "pvc", "pv", "po", "rc", "quota", "sa", "svc", "crd", "crds", "ds", "deploy", "rs", "sts", "app", "apps", "appset", "appsets", "appproj", "appprojs", "hpa", "cj", "cr", "crs", "cert", "certs", "csr", "dwoc", "dwr", "ev", "dep", "cpol", "crcr", "gr", "pol", "rcr", "ur", "amcfg", "am", "pmon", "prb", "prom", "promrule", "smon", "ruler", "ing", "netpol", "pdb", "pc", "sc", "cpolr", "polr", "dw", "dwt",
				},
			},
			"-n": Completion{
				Level:   "args",
				IsShell: true,
				Shell:   "kubectl get ns|sed 1d",
			},
			"-l": Completion{
				Level:   "show labels",
				IsShell: true,
				Shell:   "kubectl get po --show-labels|sed 1d|awk '{print $6}'",
			},
		},
	},
	"kill": Completion{
		Level:   "kill",
		IsShell: true,
		Shell:   "ps -ef|sed 1d",
	},
}

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
	// fmt.Println(result)
	// fmt.Printf("in %s %v %d\n", in, tmp, len(tmp))
	// switch tmp[0] {
	// case "kubectl", "demo":
	// 	if len(tmp) == 1 {
	// 		// os.Stdout.Write([]byte("get describe edit set"))
	// 		for _, x := range strings.Split("get describe edit set", " ") {
	// 			fmt.Println(x)
	// 		}
	// 		return
	// 	} else if len(tmp) > 1 {
	// 		switch tmp[len(tmp)-1] {
	// 		case "-n":
	// 			for _, x := range strings.Split("default kube-system eclipse-che", " ") {
	// 				fmt.Println(x)
	// 			}
	// 		case "describe", "edit", "get":
	// 			for _, x := range strings.Split("po pod ns namespace deploy job ingress svc cm secret", " ") {
	// 				fmt.Println(x)
	// 			}
	// 		}
	// 	}
	// 	if tmp[len(tmp)-1] == "-n" {
	// 		execCmd("ps -ef")
	// 	}
	// default:
	// 	execCmd(in)
	// }

	if value, ok := Completes[result[0]]; ok {
		// 判断一级命令
		if len(result) == 1 {
			if value.IsShell {
				// fmt.Println("one")
				execCmd(value.Shell)
			} else {
				// fmt.Println("two")
				fmt.Println(strings.Join(value.Cmd, "\n"))
				fmt.Println(strings.Join(value.Args, "\n"))
			}
		} else {
			// 判断一级后面的最后一个命令
			// fmt.Println("three")
			if value_daughter, ok := value.Daughter[result[len(result)-1]]; ok {
				if value_daughter.IsShell {
					// fmt.Println("four")
					execCmd(value_daughter.Shell)
				} else {
					// fmt.Println("five")
					if value_daughter.Cmd != nil {
						fmt.Println(strings.Join(value_daughter.Cmd, "\n"))
					}
					if value_daughter.Args != nil {
						fmt.Println(strings.Join(value_daughter.Args, "\n"))
					}
				}
			} else {
				// fmt.Println("six")
				// 根据命令长度智能补全命令
				if len(result) == 2 {
					// fmt.Println("7")
					t_one := []string{}
					for _, k := range value.Cmd {
						if strings.HasPrefix(k, result[1]) {
							t_one = append(t_one, strings.Replace(k, result[1], "", 1))
						}
					}
					fmt.Printf(strings.Join(t_one, "\n"))
				} else if len(result) == 3 && !isLastWorkSpace {
					// 补全命令
					t_two := []string{}
					if value.Daughter["get"].IsShell {
						// fmt.Println("9", t_two)
						rs_string, err := execCmdString(value.Daughter["get"].Shell)
						if err != nil {
							// fmt.Println("9,1 ", err.Error())
							fmt.Println(err.Error())
							return
						}
						// fmt.Println("9.2", t_two)
						// 补全 -n中提供的数据后缀
						for _, v := range strings.Split(rs_string, "\n") {
							if strings.HasPrefix(v, result[len(result)-1]) {
								t_two = append(t_two, strings.Replace(v, result[len(result)-1], "", 1))
							}
						}
						// fmt.Println("9.3", t_two)
					} else {
						// fmt.Println("10")
						for _, k := range value.Daughter["get"].Cmd {
							if strings.HasPrefix(k, strings.TrimSpace(result[len(result)-1])) {
								t_rs := strings.Replace(k, strings.TrimSpace(result[len(result)-1]), "", 1)
								if t_rs != "" && t_rs != " " {
									t_two = append(t_two, t_rs)
								}
							}
						}
					}
					fmt.Printf(strings.Join(t_two, "\n"))
				} else {
					// fmt.Println("8")
					// 补全数据
					t_two := []string{}
					// 优先匹配value.Daughter里面的值
					// for _, k := range value.Daughter["get"].Cmd {
					// 	if strings.HasPrefix(k, strings.TrimSpace(result[len(result)-1])) {
					// 		t_rs := strings.Replace(k, strings.TrimSpace(result[len(result)-1]), "", 1)
					// 		if t_rs != "" && t_rs != " " {
					// 			t_two = append(t_two, t_rs)
					// 		}
					// 	}
					// }

					// 补全上一个命令的结果，查询并替换已有数据
					// 没有发现命令则获取上一步的命令 获取资源后进行资源匹配
					if len(t_two) == 0 && isLastWorkSpace {
						// 补全实时数据结果
						// fmt.Printf("11 kubectl get %s -A\n", result[len(result)-1])
						// execCmd(fmt.Sprintf("kubectl get %s -A", result[len(result)-1]))

						// 补全数据 有空格
						// 如果获取最后一个参数无数据 则执行该命令

						rs_string, err := execCmdString(fmt.Sprintf("kubectl get %s -A", result[len(result)-1]))
						if err != nil {
							// fmt.Println("9,1 ", err.Error())
							// 执行命令错误时执行整个命令
							if strings.Contains(err.Error(), "exit") {
								// fmt.Println("12 ", in, result)
								if !strings.Contains(in, "edit") && result[1] != "edit" {
									execCmd(in)
								} else {
									execCmd(strings.Replace(in, "edit", "get", 1))
								}
							} else {
								fmt.Println(err.Error())
							}
							return
						} else {
							// fmt.Println("9.2", t_two)
							// 补全 -n中提供的数据后缀
							for _, v := range strings.Split(rs_string, "\n") {
								if strings.TrimSpace(v) != "" {
									t_two = append(t_two, v)
								}
							}
						}

					} else if len(t_two) == 0 && !isLastWorkSpace {
						// 补全命令
						// fmt.Println("12")
						// for _, v := range value.Daughter["get"].Cmd {
						// 	if strings.HasPrefix(v, result[len(result)-1]) {
						// 		t_two = append(t_two, strings.Replace(v, result[len(result)-1], "", 1))
						// 	}
						// }

						// 补全命令
						if value.Daughter["get"].IsShell {
							// fmt.Println("13")
							// fmt.Println("9", t_two)
							rs_string, err := execCmdString(value.Daughter["get"].Shell)
							if err != nil {
								// fmt.Println("9,1 ", err.Error())
								fmt.Println(err.Error())
								return
							}
							// fmt.Println("9.2", t_two)
							// 补全 -n中提供的数据后缀
							for _, v := range strings.Split(rs_string, "\n") {
								if strings.HasPrefix(v, result[len(result)-1]) {
									t_two = append(t_two, strings.Replace(v, result[len(result)-1], "", 1))
								}
							}
							// fmt.Println("9.3", t_two)
						} else {
							// fmt.Println("14")
							// fmt.Println("10")
							for _, k := range value.Daughter["get"].Cmd {
								if strings.HasPrefix(k, strings.TrimSpace(result[len(result)-1])) {
									t_rs := strings.Replace(k, strings.TrimSpace(result[len(result)-1]), "", 1)
									if t_rs != "" && t_rs != " " {
										t_two = append(t_two, t_rs)
									}
								}
							}
						}

						// 补全可能缺失的数据
						if value_maybe, ok := value.Daughter[result[len(result)-2]]; ok {
							if value_maybe.IsShell {
								// 补全实时数据结果
								// execCmd(value_maybe.Shell)
								rs_string, err := execCmdString(value_maybe.Shell)
								if err != nil {
									// fmt.Println("9,1 ", err.Error())
									fmt.Println(err.Error())
									return
								}
								// fmt.Println("9.2", t_two)
								// 补全 -n中提供的数据后缀
								for _, v := range strings.Split(rs_string, "\n") {
									if strings.HasPrefix(v, result[len(result)-1]) {
										t_two = append(t_two, strings.Replace(v, result[len(result)-1], "", 1))
									}
								}
							} else {
								// fmt.Println("10")
								for _, k := range value_maybe.Cmd {
									if strings.HasPrefix(k, strings.TrimSpace(result[len(result)-1])) {
										t_rs := strings.Replace(k, strings.TrimSpace(result[len(result)-1]), "", 1)
										t_two = append(t_two, t_rs)
									}
								}
							}
						}
						// fmt.Println("15")
						// 补全命令无效 获取上级命令的结果 并补全数据prefix数据
						if len(t_two) == 0 {
							if value_daughter2, ok := value.Daughter[result[len(result)-2]]; ok {
								if value_daughter2.IsShell {
									// 补全实时数据结果
									// fmt.Printf("13 %s\n", value_daughter2.Shell)
									rs_string, err := execCmdString(value_daughter2.Shell)
									if err != nil {
										fmt.Println(err.Error())
										return
									}
									// 补全 -n中提供的数据后缀
									for _, v := range strings.Split(rs_string, "\n") {
										if strings.HasPrefix(v, result[len(result)-1]) {
											t_two = append(t_two, strings.Replace(v, result[len(result)-1], "", 1))
										}
									}
								} else {
									// 补全cmd命令后缀
									// fmt.Printf("14 %s\n", value_daughter2.Shell)
									for _, v := range value_daughter2.Cmd {
										if strings.HasPrefix(v, result[len(result)-2]) {
											t_two = append(t_two, strings.Replace(v, result[len(result)-2], "", 1))
										}
									}
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
		fmt.Println("7")
		execCmd(in)
	}
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
	Short: "A brief description of your application",
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
	// if len(os.Args) == 2 {
	// 	if os.Args[1] == "completion" {
	// 		err := rootCmd.Execute()
	// 		if err != nil {
	// 			os.Exit(1)
	// 		}
	// 	} else {
	// 		parseCmd(strings.Join(os.Args[1:], " "))
	// 	}
	// } else if len(os.Args) == 1 {
	// 	err := rootCmd.Execute()
	// 	if err != nil {
	// 		os.Exit(1)
	// 	}
	// } else {
	// 	parseCmd(strings.Join(os.Args[1:], " "))
	// }

	if len(os.Args) == 2 && os.Args[1] == "completion" {
		err := rootCmd.Execute()
		if err != nil {
			os.Exit(1)
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
