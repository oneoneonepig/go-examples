FROM golang:1.14 AS builder
ADD . /go/src/metrics-api
WORKDIR /go/src/metrics-api
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /metrics-api .

FROM alpine:latest AS runtime
ENV PORT 8080
RUN apk --no-cache add ca-certificates
COPY --from=builder /metrics-api ./
RUN chmod +x ./metrics-api
ENTRYPOINT ["./metrics-api"]
