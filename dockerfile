FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

RUN go build -o /todolist ./cmd/app/main.go

EXPOSE 9090

CMD ["/todolist"]