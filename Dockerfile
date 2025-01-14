FROM golang:1.23


LABEL maintainer="Mohsen taheri (mhthrh@gmail.com)"
LABEL version="v1.0.1"
LABEL description="basic go image with simple API"

RUN apt update && apt install vim -y

ENV APP_ENV=test
ENV PORT=8585
ENV IP="0.0.0.0"

WORKDIR app

COPY . .
RUN ["/bin/bash", "-c", "ls -al"]
RUN go mod download && go mod verify

COPY . .
RUN go build -o api cmd/main.go

CMD ["./api"]