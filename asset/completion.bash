__smkubectl-cli_debug()
{
    local file="$BASH_COMP_DEBUG_FILE"
    if [[ -n ${file} ]]; then
        echo "$*" >> "${file}"
    fi
}

_k() {
  local matches namespace result cmd trigger cur
  # cmd="${COMP_WORDS[0]}"
  # if [[ $cmd == \\* ]]; then
  #   cmd="${cmd:1}"
  # fi
  # cmd="${cmd//[^A-Za-z0-9_=]/_}"
  # COMPREPLY=()
  # trigger=${FZF_COMPLETION_TRIGGER-'**'}
  # cur="${COMP_WORDS[COMP_CWORD]}"

  # echo "|"
  # echo $*
  # echo $@
  # echo $#
  # echo $0
  # echo $1
  # echo ${cmd}
  # echo ${trigger}
  # echo ${cur}
  # echo ${COMP_WORDS}
  # echo ${COMP_WORDS[COMP_CWORD]}
  # echo ${COMP_WORDS[COMP_CWORD-1]}
  # echo ${COMP_LINE}
  # echo ${LBUFFER}
  # echo "|"

  # 获取一次性结果
  result=$(command smkubectl smart "${COMP_LINE}")

  __smkubectl-cli_debug "result ${result}"
  namespace=$(command echo "${result}"|head -1|awk '/^NAMESPACE/ {print "yes"}')
  __smkubectl-cli_debug "namespace ${namespace}"
  # 根据标题是否含有NAMESPACE动态切换显示结果
  if [[ -n "$namespace" ]]; then
    matches=$(command echo "${result}" | FZF_DEFAULT_OPTS="--height ${FZF_TMUX_HEIGHT:-50%} --header-lines=1 --min-height 15 --reverse $FZF_DEFAULT_OPTS --preview 'echo {}' --preview-window down:3:wrap $FZF_COMPLETION_OPTS" smkubectl -m|awk '{print "-n "$1" "$2}'|tr '\n' ' ')
  else
    matches=$(command echo "${result}" | FZF_DEFAULT_OPTS="--height ${FZF_TMUX_HEIGHT:-50%} --header-lines=1 --min-height 15 --reverse $FZF_DEFAULT_OPTS --preview 'echo {}' --preview-window down:3:wrap $FZF_COMPLETION_OPTS" smkubectl -m|awk '{print $1}'|tr '\n' ' ')
  fi
  __smkubectl-cli_debug "matches ${matches}" 
  #_describe 'command' ns
  if [[ -n "$matches" ]]; then
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"
    COMPREPLY=( $(compgen -W "$matches" -- "${cur}" ) )
    COMP_LINE+="$matches"
    # COMP_WORDS+=(${matches})
    # echo "Updated COMP_WORDS:" "${COMP_WORDS[@]}"
    # compgen -W "${COMP_WORDS[@]}" -- ${cur}
    # echo -e "\r${COMP_LINE}${matches}"
    # printf "\r${COMP_LINE}"
    echo -e "\r${COMP_LINE}"
    # echo -n "\r${COMP_LINE}"
    # printf "${COMPREPLY[*]}"
    # 更新命令行数据
    # 更新命令行数据
    eval "history -s \"${COMP_LINE}\""
    eval ${COMP_LINE}
  fi
  #COMPREPLY=()
  # clear
}


if [[ $(type -t compopt) = "builtin" ]]; then
  complete -o default -F _k k git kubectl go
else
  complete -o default -o nospace -F _k k git kubectl go
fi
