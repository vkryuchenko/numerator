FROM golang:alpine AS builder

ENV GOPATH="/go:/app"
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go build -o numerator src/numerator.go

FROM alpine:latest
LABEL maintainer="Vyacheslav Kryuchenko <slavyan.85@mail.ru>"
#RUN apk add --no-cache ca-certificates
COPY --from=builder /app/numerator /bin/numerator
VOLUME /data
EXPOSE 8888
CMD ["/bin/numerator"]