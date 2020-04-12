# godock

[![Go Report Card](https://goreportcard.com/badge/github.com/nukesz/godock)](https://goreportcard.com/report/github.com/nukesz/godock)

Build docker images without creating a dockerfile.

## Usage

```sh
# Get godock
go get github.com/nukesz/godock

# Use it as you would the docker command without the -f flag
godock build -t myname/myapp .
```

