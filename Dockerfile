#build stage
FROM golang:1.14-alpine AS builder
WORKDIR /go/src/GIG-Scripts
COPY . .
RUN apk add --no-cache git
RUN git clone https://github.com/LSFLK/GIG-SDK.git /go/src/GIG-SDK
RUN go get github.com/revel/revel
RUN go get github.com/PuerkitoBio/goquery
RUN go get github.com/pkg/errors
RUN go get github.com/unidoc/unidoc
RUN go get golang.org/x/net/html
RUN go get golang.org/x/image/tiff/lzw
RUN go get gopkg.in/mgo.v2/bson
RUN go build /go/src/GIG-Scripts/kavuda/crawl.go

#running stage
FROM alpine:3.9
COPY --from=builder /go/src/GIG-Scripts/kavuda/crawl /app/GIG-Scripts/kavuda/crawl
RUN ls -l /app/GIG-Scripts/kavuda/
CMD ./app/GIG-Scripts/kavuda/crawl