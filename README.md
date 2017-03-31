# Golang Basics Workshop and Testing

Learning New Programming Languages First Workshop on Golang

Read the [Golang GitBook](https://jbelmont.github.io/golang-basics-and-testing/)

### Install Go

**[GO](https://golang.org/doc/install)**

### GOPATH
1. Define GOPATH
    1. Using environment variables set `GOPATH=/Users/marcelbelmont/go`
    2. This is an example of MAC OS X
    3. Windows users might have to edit the environment variables through advanced settings
    4. Although the `MSI` installer should do this

### Workspace

A workspace is a directory hierarchy with three directories at its root:

* src contains Go source files
* pkg contains package objects
* bin contains executable commands

* The GOPATH environment variable specifies the location of your workspace.
* It defaults to a directory named go inside your home directory, so $HOME/go on Unix
* %USERPROFILE%\go (usually C:\Users\YourName\go) on Windows.

Create your projects under the workspace

* My path is `/Users/marcelbelmont/go` and I place all my projects under

** `/Users/marcelbelmont/go/src/github.com/jbelmont` **

This is a convention followed in GO to make your project `go get` able

`MAC` users should run the following command `mkdir -p ~/go/src/github.com/${github-username}` in terminal

Optionally just open finder and right click and folder structure manually
Windows users can do the same thing

### Editor

I would recommend the `VSCODE Text Editor` it has a nice Go extension

Download [CODE](https://code.visualstudio.com/)

Install Go extension by clicking extension icon and type go in the market place input box then install it

* The extension will prompt you to install some missing packages you should do this 
* Lint, Formatting, and more will be done by the EDITOR