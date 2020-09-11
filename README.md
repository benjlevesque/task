# task CLI

A simple task CLI written in GO. 

> This project is a personal exploration of GO, inspired from [gophercises](https://gophercises.com/).


## Install

```
curl -sfL https://raw.githubusercontent.com/benjlevesque/task/master/install.sh | sh
./bin/task
```


## Build

```
git clone https://github.com/benjlevesque/task $GOPATH/src/github.com/benjlevesque/task
cd $GOPATH/src/github.com/benjlevesque/task
go get ./...
go install .
# optional. Adds autocompletion for oh-my-zsh. See task help completion 
task completion zsh > ~/.oh-my-zsh/completions/_task
task
``` 
