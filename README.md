## Install
1. download && install from https://go.dev/doc/install or using homebrew with `brew install go`
2. confirm _/usr/local/go/bin_ added to PATH

## Create
- create a repo, clone and cd into it
- `go mod init github.com/{USERNAME}/{REPO_NAME}`
- `mkdir -p cmd/{REPO_NAME} && touch cmd/{REPO_NAME}/main.go`

## Run
- `go run cmd/{REPO_NAME}/main.go`