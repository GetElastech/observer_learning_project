FROM golang:1.17 as build

COPY dep /app/dep
WORKDIR /app/dep
RUN go mod download

COPY src/go.mod src/go.sum /app/src/
WORKDIR /app/src
RUN go mod download

COPY src /app/src

WORKDIR /app/dep
RUN apt update; apt install -y cmake
RUN find $GOPATH
# RUN make install-tools
# RUN go mod download github.com/onflow/flow-go/crypto@v0.24.3
# RUN cd $GOPATH/pkg/mod/github.com/onflow/flow-go/crypto@v0.24.3 && go generate

WORKDIR /app/src
# RUN make install-tools
RUN go mod download
RUN go mod tidy
RUN go mod download github.com/onflow/flow-go/crypto@v0.24.3
RUN ls $GOPATH/pkg/mod/github.com/onflow/flow-go/crypto@v0.24.3
RUN cd $GOPATH/pkg/mod/github.com/onflow/flow-go/crypto@v0.24.3 && go generate && go build
RUN go generate -v
RUN go build -tags=relic -v -o /app ./...

## (3) Add the statically linked binary to a distroless image
FROM ubuntu as production

COPY --from=build /app/m /observer

CMD ["/observer"]
