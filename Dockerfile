FROM golang:1.23.0-alpine as build-stage

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/

RUN go build -o server .

FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /app/scripts/build/server.json .
RUN ls -a
COPY --from=build-stage /app/cmd/server .

EXPOSE 5252

CMD ["./server"]