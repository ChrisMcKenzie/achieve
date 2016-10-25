# Achieve

A modern tool for development task automation. `Achieve` aims to provide a 
modern way to perform common task in development without having to shoehorn 
them into a `Makefile` which really wasn't meant for non-file manipulating tasks

## Getting Started

Achieve is built using `go 1.7+` and does not require any other dependencies 
past that.

### Installing

```
go get github.com/chrismckenzie/achieve
```

To run a test task you can do the following

```
cd $GOPATH/src/github.com/chrismckenzie/achieve

achieve build
```

this should output a compiled achieve binary in the achieve directory
