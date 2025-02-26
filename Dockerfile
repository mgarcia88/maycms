# Use the official Go image based on Alpine
FROM golang:1.24.0-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY framework /app/
COPY domain /app/
COPY go.mod /app/
COPY go.sum /app/

# Download dependencies
RUN go build -o /server

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /server /server

EXPOSE 8080

ENTRYPOINT [ "/server" ]