# go_league_cli
This is a cli tool that take in a file of games played and computes the league standings

![Release Status](https://github.com/kbnox/go_league_cli/actions/workflows/build_and_test.yml/badge.svg)

![Release Status](https://github.com/kbnox/go_league_cli/actions/workflows/release.yml/badge.svg)


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