set fish_plugins autojump
source /etc/profile.d/autojump.fish
set AUTOJUMP_IGNORE_CASE = 1

function fish_user_key_bindings
    # Alt+l : print the file contents of the current folder.
    bind \el 'echo; and ls; and commandline -f repaint'

    # Alt+a : print the file (hidden included) contents of the current folder.
    bind \ea 'echo; and ls -a; and commandline -f repaint'

    # Alt+. : insert last argument from the last command.
    bind \e. 'history-token-search-backward'

    # Alt+o : pipe to xin
    bind \eo 'commandline -a "| xin;"'
end

set fish_greeting ""

# Serve folder on port 8000.
alias serve='/bin/sh -c "(cd $argv[1] && python -m http.server)"'

# Aliases
alias -='cd -'
alias ...='cd ../..'
alias .... 'cd ../../..'
alias ..... 'cd ../../../..'
alias ...... 'cd ../../../../..'

# Calendar start on monday.
alias cal="cal -m"

# Better diff.
alias diff='grc diff -u'

# Colorized df
alias df='grc df'

# File sizes in current directory.
alias du='du -h --summarize * | sort -h'

# Colored grep.
alias grep='grep --color=auto'

# Colored ls, (-F) append '/' to directories, (-X) sort alphabetically, (-v)
# natural sort of numbers.
alias ls='ls --color=auto -F -X -v'

# stdin to clipboard.
alias xin='xclip -in -selection clip'
alias xout='xclip -out'

# Cleaner output
alias time='time -p'
