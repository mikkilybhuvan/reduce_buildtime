# syntax=docker/dockerfile:1.7-labs
FROM golang:1.22 AS build
WORKDIR /src

# 1) Cache deps
COPY go.mod .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# 2) Copy code and build
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/myapp ./cmd/myapp

# 3) Minimal runtime image
FROM gcr.io/distroless/base-debian12
COPY --from=build /out/myapp /myapp
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/myapp"]