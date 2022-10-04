FROM golang:1.18 as build-env

WORKDIR $GOPATH/src
COPY . .

ENV GO111MODULE=auto

RUN go mod tidy
RUN go build -o /go/bin/service -v cmd/main.go

FROM gcr.io/distroless/base
COPY --from=build-env --chown="nonroot:nonroot" /go/bin/service /
USER nonroot

COPY internal/infra/database/migration internal/infra/database/migration

CMD ["/service"]
