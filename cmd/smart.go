/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"strings"

	c "github.com/lflxp/smkubectl/completion"
	g "github.com/lflxp/smkubectl/completion/git"
	"github.com/lflxp/smkubectl/completion/kubectl"
	"github.com/spf13/cobra"
)

var debugLevel bool

// zsh kubectl api-resources固定配置
var fast bool

// smartCmd represents the smart command
var smartCmd = &cobra.Command{
	Use:   "smart",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 动态调整日志输出级别
		// if debugLevel {
		// 	lvl.Set(slog.LevelDebug)
		// }
		// slog.Debug("smart cmd", slog.Any("args", args))
		// ParseCmd(strings.Join(args, " "))

		cd := &c.Command{
			Raw: strings.Join(args, " "),
		}

		tmp := strings.Split(args[0], " ")
		switch tmp[0] {
		case "k", "kubectl":
			kubectl.Start(cd)
		case "g", "git":
			g.Start(cd)
		default:
			fmt.Println("ERROR")
			slog.Error("命令不存在", "CMD", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(smartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// smartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// smartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	smartCmd.Flags().BoolVarP(&debugLevel, "debug", "d", false, "Log Level")
	smartCmd.Flags().BoolVarP(&fast, "fast", "f", false, "是否开启快速api-resources对齐")
}
