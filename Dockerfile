FROM golang:alpine
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN apk add --no-cache git mercurial
RUN go get -d ./...
RUN go build -o main .
CMD ["./main"]