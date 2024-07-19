# Build
FROM golang:1.22 AS build
WORKDIR /app
COPY . .
# Strip debug symbols from the compiled binary for a smaller image and faster startups
ENV GO_FLAGS="-ldflags '-w -s'"
RUN make compile

# Runtime
FROM alpine:3
WORKDIR /app
COPY --from=build /app/dist/helloserver .
ENTRYPOINT ["/app/helloserver"]
