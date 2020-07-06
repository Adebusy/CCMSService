FROM golang:1.14

WORKDIR /go/src/CCMSService
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8060
CMD ["CCMSService"]