# === [ ~/.bashrc ] ============================================================

# green prompt for regular users
PS1='\[\e[0;32m\][\w]\$\[\e[0m\] '

alias ls='ls --color=auto -a -F'
alias grep='grep --color=auto'

export EDITOR="geany -i"
export GOPATH=/home/u/goget:/home/u/Desktop/go
export PATH=/home/u/go/bin:/home/u/go/bin/tool:/home/u/Desktop/go/bin:/home/u/goget/bin:$PATH