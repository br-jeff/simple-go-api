FROM golang:1.20.12

RUN apt update -y && apt install -y libhyperscan-dev

WORKDIR /usr/src/app


COPY . .

RUN go build -o /usr/local/bin/app
EXPOSE 8080

CMD [ "app" ]