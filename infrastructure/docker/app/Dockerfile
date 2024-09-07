# syntax=docker/dockerfile:1

FROM golang:1.22

# Set destination for COPY
WORKDIR /var/www

ENV GOPATH=/var/www/.go

# Download Go modules
COPY ./src/go.mod ./src/go.sum ./
RUN go mod download

# Copy the source code
COPY ./src ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /organization ./cmd/api

# Expose port 9000
EXPOSE 9000

# Run
CMD [ "/organization" ]
