### Go installation 
- location : https://golang.org/doc/install?download=go1.11.5.linux-amd64.tar.gz
- export GOPATH=$HOME/go
- export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
- mkdir -p $HOME/go{bin,src}
- apt-get install vim
- curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
- vi ~/.vimrc
````
call plug#begin('~/.vim/plugged')

Plug 'fatih/vim-go'

call plug#end()

filetype off
filetype plugin indent on

set number
set noswapfile
set noshowmode
set ts=2 sw=2 sts=2 et
set backspace=indent,elo,start

" MAP <leader> to comma
let mapleader=","

if has("autocmd")
  autocmd FileType go set ts=2 sw=2 sts=2 noet nolist autowrite
endif
````
