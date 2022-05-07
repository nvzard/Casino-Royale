FROM golang:1.18-alpine

RUN apk add git build-base curl
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
RUN export PATH=$PATH:/casino-royale/bin

WORKDIR /casino-royale

COPY go.mod .
RUN go mod tidy
RUN go mod download

COPY . .
RUN go build -o ./bin/casino-royale .

EXPOSE 8080

CMD ["./bin/casino-royale"]
