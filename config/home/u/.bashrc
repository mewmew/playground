# === [ ~/.bashrc ] ============================================================

# green prompt for regular users
PS1='\[\e[0;32m\][\w]\$\[\e[0m\] '

alias ls='ls --color=auto -a -F -X'
alias grep='grep --color=auto'
alias diff='diff -u'
alias cal='cal -m'
alias xin='xclip -in -selection clip'

export EDITOR="geany -i"
export GOPATH=/home/u/goget:/home/u/Desktop/go
export PATH=/home/u/go/bin:/home/u/go/pkg/tool/linux_amd64:/home/u/Desktop/go/bin:/home/u/goget/bin:$PATH
