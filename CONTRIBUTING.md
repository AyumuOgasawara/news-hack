# Contributing

## On MacOS

Install go:

    brew install go

Then add to path by adding the following to `~/.zshrc`:

    export GOPATH=$HOME/go
    export PATH=$PATH:/usr/local/go/bin
    export PATH=$PATH:$GOPATH/bin

Make sure this file is sourced:

    . ~/.zshrc

Install current pre-commit dependencies:

    brew install golangci-lint
    go install golang.org/x/tools/cmd/goimports@latest

Install pre-commit:

    brew install pre-commit
