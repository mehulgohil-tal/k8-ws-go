FROM golang:1.16

WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

COPY . ./
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /inventory-service
# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 3000

# Run
CMD ["/inventory-service"]