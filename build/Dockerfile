FROM golang:1.20.5-bullseye AS builder
RUN apt-get update && apt-get upgrade -y
WORKDIR /build
COPY . .
RUN go clean --modcache \
    && go build -o rest-server cmd/rest/main.go

FROM debian:bullseye-slim
RUN apt-get update \
	&& apt-get upgrade -y \
	&& useradd -ms /bin/bash -u 1000 bfs \
	&& apt-get clean \
	&& rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /build/rest-server ./
COPY --from=builder /build/web/agenda/dist/ ./web
COPY --from=builder /build/scripts/migrations/ ./migrations
RUN mkdir -p data /app/web/documents && chown bfs /app/web/documents
USER bfs
CMD ["./rest-server"]