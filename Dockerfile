FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/github.com/Project-Prismatica/prismatica_integration_nessus
COPY . /usr/local/go/src/github.com/Project-Prismatica/prismatica_integration_nessus
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/prismatica_integration_nessus ./prismatica_integration_nessus


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/github.com/Project-Prismatica/prismatica_integration_nessus/build/prismatica_integration_nessus /bin/prismatica_integration_nessus
CMD ["prismatica_integration_nessus", "up"]
