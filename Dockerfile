FROM golang:1.17-alpine as build

WORKDIR apps/iotube-validator-service

RUN apk add --no-cache make gcc musl-dev linux-headers git protobuf protoc

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN mkdir -p $GOPATH/pkg/linux_amd64/github.com/iotexproject/ && make build

FROM alpine:latest

RUN apk add --no-cache ca-certificates
RUN mkdir -p /etc/iotube/

COPY docker/conf/api.yaml /etc/iotube/api.yaml
COPY docker/conf/monitor.yaml /etc/iotube/monitor.yaml
COPY docker/conf/signer.yaml /etc/iotube/signer.yaml

COPY --from=build /go/apps/iotube-validator-service/bin/relayer /usr/bin/relayer
COPY --from=build /go/apps/iotube-validator-service/bin/validator /usr/bin/validator

CMD [ "validator" ]