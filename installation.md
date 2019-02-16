### Go installation 
- Online playground : https://play.golang.org/
- location : https://golang.org/doc/install?download=go1.11.5.linux-amd64.tar.gz
- export GOPATH=$HOME/go
- export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
- mkdir -p $HOME/go{bin,src}
- apt-get install vim
- apt-get install golang-go
- apt install gccgo-go
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
set backspace=indent,eol,start

" MAP <leader> to comma
let mapleader=","

if has("autocmd")
  autocmd FileType go set ts=2 sw=2 sts=2 noet nolist autowrite
endif
````
- vim +PlugInstall +qa
- vim +GoInstallBinaries +qa
````
cat ~/go/src/hello_world# cat main.go 
package main

import ""fmt"


func main() {
  fmt.Println(" Hello Chaitanya")
  }
  ````
````  
#~/go/src/hello_world# go run main.go
 Hellow Chaitanya

#go build  main.go
root@chaitanya:~/go/src/hello_world# ls
main  main.go
root@chaitanya:~/go/src/hello_world# ./main 
 Hellow Chaitanya
````
-  Multiple line comments : /*     */
-  single line comments : //
-  https://golang.org/pkg/     "packages for Go lang"
-  https://golang.org/pkg/fmt/ "FMT Package doc"
-  https://golang.org/pkg/fmt/#Println  "Examples "

- /go/src# cat learn_strings.go 
````
package main

import "fmt"

func main() {
	fmt.Println("Simple String")
	fmt.Println(` 
This is a multi-line \n

string
`)
	fmt.Println(" emojis ")
	fmt.Println("\u22272")
}
````
````

root@chaitanya:~/go/src# cat learn_strings.go 
package main

import "fmt"

func main() {
	fmt.Println("Simple String")
	fmt.Println(` 
This is a multi-line \n

string
`)
	fmt.Println("vura chaitanya")
	fmt.Println('v')
	//	fmt.Println('chaitanya')
	fmt.Println("\u22272")
}
````

````
root@chaitanya:~/go/src# go run learn_strings.go 
Simple String
 
This is a multi-line \n

string

vura chaitanya
118
∧2
````

