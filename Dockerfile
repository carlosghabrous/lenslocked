FROM golang:1.23 

WORKDIR /go/src/app
COPY . .
RUN go build -o /go/bin/app/

EXPOSE 3000
CMD ["/go/bin/app/lenslocked"]