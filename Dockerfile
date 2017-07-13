FROM golang
RUN mkdir -p /go/src/lab1
WORKDIR /go/src/lab1
COPY . /go/src/lab1
RUN go-wrapper download
RUN go-wrapper install
CMD ["go-wrapper", "run"]