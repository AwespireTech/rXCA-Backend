FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=builder /app/main /app/
EXPOSE 8080
WORKDIR /app
ENTRYPOINT ["./main"]

