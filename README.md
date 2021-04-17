# oto-showcase

An example of how to use oto:

- https://github.com/pacedotdev/oto
- https://pace.dev/blog/2020/07/27/how-code-generation-wrote-our-api-and-cli.html

## Install

```shell
# outside of go.mod
go get github.com/pacedotdev/oto
```

## Templates

### wget

```shell
mkdir templates
wget https://raw.githubusercontent.com/pacedotdev/oto/master/otohttp/templates/server.go.plush -q -O ./templates/server.go.plush
```

### curl

```shell
mkdir templates
curl -s https://raw.githubusercontent.com/pacedotdev/oto/master/otohttp/templates/server.go.plush > ./templates/server.go.plush
```

## Generate

```shell
mkdir generated

oto -template ./templates/server.go.plush \
    -out ./generated/oto.gen.go \
    -ignore Ignorer \
    -pkg generated \
    ./definitions

# optional, but useful
gofmt -w ./generated/oto.gen.go ./generated/oto.gen.go
```
