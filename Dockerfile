FROM golang:alpine as builder
RUN apk add -U --no-cache ca-certificates && \
    apk add --no-cache bash git openssh
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o next .
FROM scratch
COPY --from=builder /build/next /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /
CMD ["./next"]