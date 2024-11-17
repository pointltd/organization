FROM golang:latest AS build

WORKDIR /app
COPY ./src ./
RUN GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o organization-app ./cmd/api

FROM golang:latest AS dev
COPY --from=build /app/organization-app /organization-app

ENTRYPOINT ["/organization-app"]

FROM scratch AS prod
COPY --from=build /app/organization-app /organization-app

ENTRYPOINT ["/organization-app"]