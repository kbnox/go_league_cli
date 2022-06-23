# go_league_cli
This is a cli tool that take in a file of games played and computes the league standings

![Release Status](https://github.com/kbnox/go_league_cli/actions/workflows/build_and_test.yml/badge.svg)

![Release Status](https://github.com/kbnox/go_league_cli/actions/workflows/release.yml/badge.svg)

# Tools needed
- [Go 1.18](https://go.dev/dl/)
- [ErrCheck](https://github.com/kisielk/errcheck) - useful for catching missed error checks

## Cloning the repo and getting started
### Cloning the repo

`git clone git@github.com:kbnox/go_league_cli.git`

`cd go_league_cli`

### download our dependencies

`go mod download`

## Build and Test

To build the application:
`go build ./...`

### Testing

To run tests run the bellow command:

`go test ./...`

## Running the tool

`./main.go -file test_score.txt`


## Installing the binary
You can discover the install path by running the go list command
`go list -f '{{.Target}}'`

Add the Go install directory to your system's shell path
- on Linux / Mac
    `export PATH=$PATH:/path/to/your/install/directory`
- on Windows
    `set PATH=%PATH%;C:\path\to\your\install\directory`

As an alternative, if you already have a directory like $HOME/bin in your shell path and you'd like to install your Go programs there, you can change the install target by setting the GOBIN variable using the go env command
- on Linux / Mac
    `go env -w GOBIN=/path/to/your/bin`
- on Windows
    `go env -w GOBIN=C:\path\to\your\bin`

Once you've updated the shell path, run the following to compile and install the package
    `go install`

You should be able to run your application by simply typing it's name
```
$ go_league_cli
  -file string
        The file containing the scores for each game.
```
