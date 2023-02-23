# build stage
FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
# RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
# RUN go install github.com/rubenv/sql-migrate/...@latest
RUN go build -o main main.go

#run stage
FROM alpine 
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY migrations ./migrations


EXPOSE 1313
CMD [ "/app/main" ]