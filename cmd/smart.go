/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	c "github.com/lflxp/smkubectl/completion"
	"github.com/lflxp/smkubectl/completion/docker"
	"github.com/lflxp/smkubectl/completion/git"
	"github.com/lflxp/smkubectl/completion/kill"
	"github.com/lflxp/smkubectl/completion/kubectl"
	"github.com/spf13/cobra"
)

var debugLevel bool

var smartCmd = &cobra.Command{
	Use:   "smart",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}

		cd := &c.Command{
			Raw: strings.Join(args, " "),
		}

		tmp := strings.Split(args[0], " ")
		switch tmp[0] {
		case "k", "kubectl":
			kubectl.Start(cd)
		case "g", "git":
			git.Start(cd)
		case "docker", "d":
			docker.Start(cd)
		case "kill":
			kill.Start(cd)
		}
	},
}

func init() {
	rootCmd.AddCommand(smartCmd)
	smartCmd.Flags().BoolVarP(&debugLevel, "debug", "d", false, "Log Level")
}
