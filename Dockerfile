FROM golang:1.8

WORKDIR /go/src/app
COPY . .
COPY .env.production .env

ENV GIN_MODE=release

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

CMD ["go-wrapper", "run"] # ["app"]

# Document that the service listens on port 8080.
EXPOSE 9090
