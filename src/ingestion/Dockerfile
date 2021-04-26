FROM golang:1.13-alpine as builder
# install gcc and create build folder
RUN apk add --no-cache gcc musl-dev git
RUN  mkdir -p ./build

ADD ./src /build/
WORKDIR /build

RUN go mod vendor
RUN go build -mod=vendor -v -i  -o ingestion

FROM alpine:3.10

RUN adduser -S -D -H -h /app appuser

USER appuser

WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /build/ingestion /app/ingestion

CMD ["./ingestion"]