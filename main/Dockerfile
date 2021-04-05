FROM golang:1.16

WORKDIR /go/app

COPY . .

RUN go get

RUN go install

RUN go build

ENV DB_URL "user=postgres dbname=books password=password sslmode=disable host=localhost port=5432"

CMD ["/go/app/main"]