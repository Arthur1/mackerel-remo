FROM golang:1.21-bookworm as builder

WORKDIR /opt

COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o "mackerel-remo" ./cmd/mackerel-remo

FROM gcr.io/distroless/base-debian12:latest

COPY --from=builder /opt/mackerel-remo /mackerel-remo
CMD ["/mackerel-remo"]
