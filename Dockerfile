FROM golang:1.23-bookworm AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/build/4sq_server

FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=builder /app/pb_public /app/pb_public
COPY --from=builder /app/build/4sq_server /app/4sq_server
EXPOSE 8080
CMD ["/app/4sq_server", "serve", "--http=0.0.0.0:8080"]
