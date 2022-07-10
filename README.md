# Go-TDlib

Package on Golang for TDLib. Current version support TDLib: v1.8.0

## Install  

You can simply run `go get` this package.

```
$ go get github.com/wh3r3areyou/go-tdlib
```


## Installation

To work with this package, you need to install tdlib on your system (I haven't tested it on Windows, and I advise you to work in a docker container).

Use the installation generator [Tdlib build generator](https://tdlib.github.io/td/build.html)
Use TDLib build instructions with checkmarked Install built TDLib to /usr/local instead of placing the files to td/tdlib.
It is mandatory that the files are available in usr/local.

If hit any build errors, refer to [Tdlib build instructions](https://github.com/tdlib/td#building)


## Docker
I advise you to work through a docker container with the correct installation of tdlib. This dockerfile serves as an example, but I do not advise deleting the necessary packages, and even more so I do not advise deleting the tdlib installation.

Place this dockerfile in your project and organize the volume for project
```
FROM golang:1.18-alpine AS golang

WORKDIR /

RUN apk add --no-cache \
        ca-certificates

RUN apk add --no-cache --virtual .build-deps \
        g++ \
        make \
        cmake \
        git \
        gperf \
        libressl-dev \
        zlib-dev \
        zlib-static \
        alpine-sdk openssl-dev gperf \
        linux-headers;

RUN git clone https://github.com/tdlib/td.git && \
    cd td && \
    git checkout v1.8.0 && \
    mkdir build && \
    cd build && \
    cmake -DCMAKE_BUILD_TYPE=Release .. && \
    cmake --build .  && \
    make install

WORKDIR /app

COPY . .
```

```
$ docker build -t tdlib-docker . 
$ docker run -it tdlib-docker
```

## Thx!
https://github.com/Arman92/go-tl-parser - Generates JSON or Go structs/methods of a Telegram .tl file Adds every single comment (Structs, Struct members, Methods, Method arguments)