FROM golang:1.14-alpine as build
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build main.go

FROM alpine:3.9
RUN apk --update add ca-certificates
COPY --from=build /app/main /
ENTRYPOINT [ "/main" ]