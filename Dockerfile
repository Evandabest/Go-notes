
FROM golang:1.23.6
WORKDIR /Go

COPY main.go /Go/
COPY print.go /Go/

RUN go mod init example.com/name
RUN go build
RUN ./name

RUN useradd app
USER app

CMD ["build", "go build", "run", "./name"]

#docker run --rm -it my-golang-app /bin/sh




