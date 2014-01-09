# === [ ~/.bashrc ] ============================================================

# green prompt for regular users
PS1='\[\e[0;32m\][\w]\$\[\e[0m\] '

alias ls='ls --color=auto -a -F'
alias grep='grep --color=auto'

export EDITOR="geany -i"
export GOPATH=/home/u/goget:/home/u/Desktop/go
export PATH=/home/u/.gem/ruby/2.0.0/bin:$PATH
export PATH=/home/u/go/bin:/home/u/go/pkg/tool/linux_amd64:/home/u/Desktop/go/bin:/home/u/goget/bin:$PATH

export SMLNJ_HOME="/usr/lib/smlnj"
export PATH=$PATH:/usr/lib/smlnj/bin
