FROM golang:latest AS build

WORKDIR /app

# Download Go modules
COPY ./src ./

RUN go mod tidy
RUN go mod vendor

RUN GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o organization-app ./cmd/api

FROM golang:latest AS dev

WORKDIR /app

COPY ./src ./

RUN go mod download
RUN go mod vendor

RUN go install github.com/air-verse/air@latest

ENTRYPOINT ["air", "-build.cmd", "go build -o ./bin/organization-app ./cmd/api", "-build.bin", "./bin/organization-app"]

FROM golang:1.22 AS prod

WORKDIR /app

# Download Go modules
COPY ./src ./

RUN go mod tidy
RUN go mod vendor

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /organization ./cmd/api

# Run
CMD [ "/organization" ]