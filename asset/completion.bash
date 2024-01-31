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

  echo "|"
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
  echo ${COMP_LINE}
  echo ${LBUFFER}
  echo "|"

  # 获取一次性结果
  result=$(command smkubectl smart -d "${COMP_LINE}")

  __smkubectl-cli_debug "result ${result}"
  namespace=$(command echo "${result}"|head -1|awk '/^NAMESPACE/ {print "yes"}')
  __smkubectl-cli_debug "namespace ${namespace}"
  # 根据标题是否含有NAMESPACE动态切换显示结果
  if [[ -n "$namespace" ]]; then
    matches=$(command echo "${result}" | FZF_DEFAULT_OPTS="--height ${FZF_TMUX_HEIGHT:-50%} --header-lines=1 --min-height 15 --reverse $FZF_DEFAULT_OPTS --preview 'echo {}' --preview-window down:3:wrap $FZF_COMPLETION_OPTS" smkubectl -m|awk '{print "-n "$1" "$2}'|tr '\n' ' ')
  else
    matches=$(command echo "${result}" | FZF_DEFAULT_OPTS="--height ${FZF_TMUX_HEIGHT:-50%} --header-lines=1 --min-height 15 --reverse $FZF_DEFAULT_OPTS --preview 'echo {}' --preview-window down:3:wrap $FZF_COMPLETION_OPTS" smkubectl -m|awk '{print $1}'|tr '\n' ' ')
  fi
  #_describe 'command' ns
  if [[ -n "$matches" ]]; then
    # COMPREPLY=( $(compgen -W "${result}" -- ${cur}) )
    # COMPREPLY=( $(compgen -W "${COMP_LINE}" -- $matches) )
    # COMPREPLY=( $(compgen -W "${COMP_LINE}$matches" ))
    COMPREPLY=( $(compgen -W "${matches}" ${COMP_WORDS[${COMP_CWORD}]}))
    # COMPREPLY+=( $(compgen -W "$matches" ))
    # ${COMP_LINE}="${COMP_LINE}$matches"
    # local idx=${#COMPREPLY[*]}
    # while [[ $((--idx)) -ge 0 ]]; do
    #     COMPREPLY[$idx]=${COMPREPLY[$idx]#"$matches"}
    # done
    # COMP_LINE+="$matches" 
  fi
}

if [[ $(type -t compopt) = "builtin" ]]; then
    complete -o default -F _k k git kubectl go
else
    complete -o default -o nospace -F _k k git kubectl go
fi
