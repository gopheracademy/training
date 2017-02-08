name: GettingStarted
video: 
description: 
level: Beginner
topic: Go
# Getting Started With Go
## Getting Started With Go

---
name: Introduction
video: 
thumb:
github:
# Introduction
## 

Brian Ketelsen - me@brianketelsen.com

- @bketelsen 
- github.com/bketelsen
- https://brianketelsen.com
- https://gophercon.com
- https://blog.gopheracademy.com
- http://www.goinactionbook.com/---
name: Installing Go
video: 
thumb:
github:
# Installing Go
## 

Mac or Linux

Download and install Go - always use the packages from golang.org - never use homebrew or apt-get, yum, etc. They're broken, or worse -- modified by someone upstream.

Set a GOPATH in .bashrc, .bash_profile, .zshrc etc:

	export GOPATH=$HOME/go

Add go binaries (compilers and tools) to your path:

	export PATH=$PATH:/usr/local/go/bin

Log out and back in to get the changes or

	$ source .bashrc

to hot-reload.---
name: Installing Go (win)
video: 
thumb:
github:
# Installing Go 
## 

Windows

Download and install Go - Use MSI installer

Set a GOPATH in user Environment Variables

	GOPATH=%userdir%/go

Add go binaries (compilers and tools) to your path:

	C:\go\bin---
name: Verifying Installation
video: 
thumb:
github:
# Verifying Installation
## 

From a command prompt:
	
	go version

You should see something like:

	go version 1.7.3 linux/amd64---
name: Download Course Material
video: 
thumb:
github:
# Download Course Material
## 

Now that you have a Go development environment set up, download the course material for this class:

	go get -u github.com/gophertrain/material---
name: Test your setup
video: 
thumb:
github:
# Test your setup
## 

cd $GOPATH/src/github.com/gophertrain/material/gettingstarted/exercises/hello
	go run main.go

You should see:

	Hello World!

Nothing like starting with the classics :)---
name: Editing Go Code
video: 
thumb:
github:
# Editing Go Code
## 

Popular Go Editors:

- vim and neovim with vim-go plugin *Favorite terminal
- emacs with go-mode.el
- Visual Studio Code with vscode-go (works with debugging!) *Favorite GUI
- Atom with go-plus
- IntelliJ IDEA (Gogland)---
name: Homework Tonight
video: 
thumb:
github:
# Homework Tonight
## 

Pick your favorite editor and configure it for Go development.

Do not choose this week to learn VIM or Emacs, please :) It will only end in sorrow. Don't do this in class today, we're moving FAST and you'll miss important things.

https://github.com/fatih/vim-go-tutorial (vim)

http://tleyden.github.io/blog/2014/05/22/configure... (emacs)

http://marcio.io/2015/07/supercharging-atom-editor... (atom)

https://github.com/Microsoft/vscode-go (VS Code)---
name: The Go Playground
video: 
thumb:
github:
# The Go Playground
## 

Even if you don't have an editor configured locally you can still play with Go from your browser.

.link https://play.golang.org

Limitations:

The Go Playground is a web service that runs on golang.org's servers. The service receives a Go program, compiles, links, and runs the program inside a sandbox, then returns the output.---
name: Playground Limitations
video: 
thumb:
github:
# Playground Limitations
## 

There are limitations to the programs that can be run in the playground:

The playground can use most of the standard library, with some exceptions. The only communication a playground program has to the outside world is by writing to standard output and standard error.

In the playground the time begins at 2009-11-10 23:00:00 UTC (determining the significance of this date is an exercise for the reader). This makes it easier to cache programs by giving them deterministic output.

There are also limits on execution time and on CPU and memory usage.

Therefore: No file IO, nothing useful with time or dates, can't use any external packages.
---
name: The Go Playground2
video: 
thumb:
github:
# The Go Playground
## 

Even with all those limitations Go developers love the Go Playground - it's a great place to share code, even if it can't run or compile. You can enter code then click the "SHARE" button which will give you a permanent URL to that code.

Try it now with this link: 

.link https://play.golang.org/p/992fMmkkxr Hello World
---
name: The Go Command
video: 
thumb:
github:
# The Go Command
## 

All of your command line interaction with Go will be through the `go` command.

Several common commands:
	
	go build some/package
	go run main.go
	go test some/package
	go get github.com/someone/package
	go install some/package---
name: Exercise
video: 
thumb:
github:
# Exercise
## 

From the command prompt type `go` and hit return to see the various tools the `go` command implements.  Try some like:

	go env
	go list
	go version
