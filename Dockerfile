FROM node:22.15-alpine3.21 AS frontend-builder
WORKDIR /build
COPY ./static .
RUN npm ci && \
    npm run build

FROM golang:1.24-alpine3.21 AS backend-builder
WORKDIR /build
COPY . .
COPY --from=frontend-builder /build/dist /build/static/dist
RUN go build -o mailweave -ldflags="-s -w -X 'main.version=$(git describe --tags --always)'" ./cmd/

FROM alpine:3.21 AS runtime
RUN apk add curl ca-certificates
COPY README.md /README.md
COPY LICENSE /LICENSE
COPY --from=backend-builder /build/mailweave /bin/mailweave
ENTRYPOINT ["/bin/mailweave"]
