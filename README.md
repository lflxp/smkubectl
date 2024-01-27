# 参数记录

1. 开启自动补全

`
autoload -U compinit && compinit -u
`

2. completion

`
compdef _gmt gmt                                                                                                                                                                                                                                                                    # 添加 _gmt 函数
_gmt() {
  local matches lbuf
  setopt localoptions
  # matches=("${(@f)$(kubectl get ns | sed 1d | awk '{print $1}' | FZF_DEFAULT_OPTS="--height ${FZF_TMUX_HEIGHT:-50%} --min-height 15 --reverse $FZF_DEFAULT_OPTS --preview 'echo {}' --preview-window down:3:wrap $FZF_COMPLETION_OPTS" fzf -m|awk '{print $1}'|tr '\n' ' ')}")
  #matches=$(command ps -ef| sed 1d | FZF_DEFAULT_OPTS="--height ${FZF_TMUX_HEIGHT:-50%} --min-height 15 --reverse $FZF_DEFAULT_OPTS --preview 'echo {}' --preview-window down:3:wrap $FZF_COMPLETION_OPTS" fzf -m|awk '{print $2}'|tr '\n' ' ')
  matches=$(command demo ${LBUFFER} | FZF_DEFAULT_OPTS="--height ${FZF_TMUX_HEIGHT:-50%} --min-height 15 --reverse $FZF_DEFAULT_OPTS --preview 'echo {}' --preview-window down:3:wrap $FZF_COMPLETION_OPTS" fzf -m|awk '{print $1}'|tr '\n' ' ')
  #matches=$(command demo ${LBUFFER} | awk '{print $2}'|tr '\n' ' ')
  #_describe 'command' ns
  if [ -n "$matches" ]; then
    LBUFFER="${LBUFFER}$matches"
  fi
  zle reset-prompt
}

zle -N _gmt
bindkey '^E' _gmt
`