FROM golang:1.19-alpine AS build
WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY bulk ./bulk
RUN go build -o /out
EXPOSE 8000

FROM alpine
COPY --from=build /out /server
ENTRYPOINT "/server"
