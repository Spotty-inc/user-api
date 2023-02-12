FROM golang:1.18-alpine

WORKDIR /app
COPY /src ./

RUN go mod download

RUN go build -o /main

EXPOSE 10001

CMD [ "/main" ]