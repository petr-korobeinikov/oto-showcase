# oto-showcase

## Install

```shell
# outside of go.mod
go get github.com/pacedotdev/oto
```

## Templates

```shell
mkdir templates \
    && wget https://raw.githubusercontent.com/pacedotdev/oto/master/otohttp/templates/server.go.plush -q -O ./templates/server.go.plush \
    && wget https://raw.githubusercontent.com/pacedotdev/oto/master/otohttp/templates/client.js.plush -q -O ./templates/client.js.plush
```

```shell
curl -s https://raw.githubusercontent.com/pacedotdev/oto/master/otohttp/templates/server.go.plush > ./templates/server.go.plush
curl -s https://raw.githubusercontent.com/pacedotdev/oto/master/otohttp/templates/client.js.plush > ./templates/client.js.plush
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

oto -template ./templates/server.go.plush \
    -out ./generated/oto.gen.go \
    -ignore Ignorer \
    -pkg generated \
    ./definitions
```
