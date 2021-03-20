# syntax = docker/dockerfile:1.0-experimental

FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
  go mod download

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
  go build -o ssh-server .

FROM gcr.io/distroless/base-debian10

WORKDIR /app
COPY --from=build /app/ssh-server ./

ENTRYPOINT [ "/app/ssh-server" ]