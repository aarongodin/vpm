FROM golang:latest AS builder

WORKDIR /app

ARG CGO_ENABLED=0
ARG GOOS=linux

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} go build -o vpm .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/vpm .

CMD ["./vpm"]
