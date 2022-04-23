# Go binaries are standalone, so use a multi-stage build to produce smaller images.

# Use base golang image from Docker Hub
FROM golang:1.18 as build

WORKDIR /app



# Install dependencies in go.mod and go.sum
COPY go.mod go.sum ./

RUN go mod tidy 
RUN go mod download

# Copy rest of the application source code
COPY . ./

RUN rm -rf ./contrib
# Compile the application to /app.
# Skaffold passes in debug-oriented compiler flags
ARG SKAFFOLD_GO_GCFLAGS
RUN echo "Go gcflags: ${SKAFFOLD_GO_GCFLAGS}"
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -mod=readonly -v -o /app/bin/ ./cmd/rpc

RUN wget -O- https://letsencrypt.org/certs/isrg-root-x1-cross-signed.pem > /tmp/root.crt

# Now create separate deployment image
FROM gcr.io/distroless/base

# Definition of this variable is used by 'skaffold debug' to identify a golang binary.
# Default behavior - a failure prints a stack trace for the current goroutine.
# See https://golang.org/pkg/runtime/
ENV GOTRACEBACK=single

WORKDIR /app
COPY --from=build /app/bin/rpc .
COPY --from=build /tmp/root.crt /etc/root.crt


EXPOSE 8080

ENTRYPOINT ["./rpc"]


