#build stage
FROM golang:1.14 AS build-env

WORKDIR /go/src/app/

#copy to workdir path
COPY main.go .

ENV CGO_ENABLED=0

#build the go app
RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

# final stage
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app/

#copy the compilate binary for workdir
COPY --from=build-env /go/src/app .

ENTRYPOINT ["./app"]