/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// 解析命令
// 只负责执行 不负责命令补全
func parseCmd(in string) {
	tmp := strings.Split(strings.TrimSpace(in), " ")
	// fmt.Printf("in %s %v %d\n", in, tmp, len(tmp))
	switch tmp[0] {
	case "kubectl":
		if len(tmp) == 1 {
			// os.Stdout.Write([]byte("get describe edit set"))
			for _, x := range strings.Split("get describe edit set", " ") {
				fmt.Println(x)
			}
			return
		} else if len(tmp) > 1 {
			switch tmp[len(tmp)-1] {
			case "-n":
				for _, x := range strings.Split("default kube-system eclipse-che", " ") {
					fmt.Println(x)
				}
			case "describe", "edit", "get":
				for _, x := range strings.Split("po pod ns namespace deploy job ingress svc cm secret", " ") {
					fmt.Println(x)
				}
			}
		}
		if tmp[len(tmp)-1] == "-n" {
			execCmd("ps -ef")
		}
	default:
		execCmd(in)
	}
}

func execCmd(in string) error {
	cmd := exec.Command("sh", "-c", in)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout // 标准输出
	cmd.Stderr = os.Stderr // 标准错误
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
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
	if len(os.Args) == 2 {
		if os.Args[1] == "completion" {
			err := rootCmd.Execute()
			if err != nil {
				os.Exit(1)
			}
		} else {
			parseCmd(strings.Join(os.Args[1:], " "))
		}
	} else if len(os.Args) == 1 {
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
