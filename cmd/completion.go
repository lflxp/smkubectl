/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const completion = `compdef _k k                                                                                                                                                                                                                                                                    # 添加 _gmt 函数

__cobra-cli_debug()
{
    local file="$BASH_COMP_DEBUG_FILE"
    if [[ -n ${file} ]]; then
        echo "$*" >> "${file}"
    fi
}

_k() {
  local matches namespace result

  # 获取一次性结果
  result=$(command demo ${LBUFFER})

  __cobra-cli_debug "namespace ${result}"
  namespace=$(command echo ${result}|head -1|awk '/^NAMESPACE/ {print "yes"}')
  __cobra-cli_debug "namespace ${namespace}"
  # 根据标题是否含有NAMESPACE动态切换显示结果
  if [[ -n "$namespace" ]]; then
    matches=$(command echo ${result} | FZF_DEFAULT_OPTS="--height ${FZF_TMUX_HEIGHT:-50%} --min-height 15 --reverse $FZF_DEFAULT_OPTS --preview 'echo {}' --preview-window down:3:wrap $FZF_COMPLETION_OPTS" fzf -m|awk '{print "-n "$1" "$2}'|tr '\n' ' ')
  else
    matches=$(command echo ${result} | FZF_DEFAULT_OPTS="--height ${FZF_TMUX_HEIGHT:-50%} --min-height 15 --reverse $FZF_DEFAULT_OPTS --preview 'echo {}' --preview-window down:3:wrap $FZF_COMPLETION_OPTS" fzf -m|awk '{print $1}'|tr '\n' ' ')
  fi
  #_describe 'command' ns
  if [ -n "$matches" ]; then
    LBUFFER="${LBUFFER}$matches"
    # LBUFFER="${BUFFER} $matches"
  fi
  zle reset-prompt
}

zle -N _k
#定义快捷键为： [Esc] [Esc]
` + "bindkey '`' _k"

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(completion)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
