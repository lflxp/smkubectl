_my_custom_complete() {
  local prepositions="of off on to"
  COMPREPLY=( $(compgen -W "${prepositions}" -- "$2") )
  tab=$(printf '\t')
  local comp="$1"
  local longest=$2
  printf "%q" "${comp}"
}

complete -o default -F _my_custom_complete mycustomprompt
