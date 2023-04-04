# Go-TDlib

Package on Golang for TDLib. Current version support TDLib: v1.8.13 (current master in TD repo)


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
I advise you to work through a docker container with a proper tdlib installation. This dockerfile file serves as an example, I do not advise deleting the necessary packages, and even more so I do not advise deleting the tdlib copy.

Put this dockerfile in your project and organize the volume for the project
```
ARG VERSION_ALPINE=3.16
ARG VERSION_GOLANG=1.18.5

FROM golang:${VERSION_GOLANG}-alpine${VERSION_ALPINE} as golang

RUN apk update && \
    apk add --no-cache ca-certificates curl g++

RUN apk update && \
    apk add --no-cache bash ca-certificates git openssh make gcc  \
    alpine-sdk openssl-dev gperf openssh gcc libressl-dev \
    zlib-dev \
    zlib-static \
    linux-headers \
    cmake && \
    rm -fr /var/cache/apk/*


COPY --from=wcsiu/tdlib:1.8.0-alpine /usr/local/include/td /usr/local/include/td
COPY --from=wcsiu/tdlib:1.8.0-alpine /usr/local/lib/libtd* /usr/local/lib/
COPY --from=wcsiu/tdlib:1.8.0-alpine /usr/lib/libssl.a /usr/local/lib/libssl.a
COPY --from=wcsiu/tdlib:1.8.0-alpine /usr/lib/libcrypto.a /usr/local/lib/libcrypto.a
COPY --from=wcsiu/tdlib:1.8.0-alpine /lib/libz.a /usr/local/lib/libz.a

WORKDIR /app

COPY . .
```

```
$ docker build -t tdlib-docker . 
$ docker run -it tdlib-docker
```

## Thx!
https://github.com/Arman92/go-tdlib - Main reference