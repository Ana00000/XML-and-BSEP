FROM golang:1.16.3-buster AS builder

# Copy core library
RUN mkdir go/settings-service
COPY settings-service/ go/settings-service

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary
RUN go build -o /app/user-service
EXPOSE 8091
CMD ["./user-service"]