# Use the official Golang image to create a build artifact.
# This is the first stage of a multi-stage build.
FROM golang:1.21.7 as builder

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod ./
# Ensure modules are downloaded and cached.
RUN go mod download

# Copy the full source code to the container
COPY . .

# Build the Go app as a static binary.
# -o specifies the output binary name, here it is "revenueCalculator".
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o revenueCalculator .

# Use a Docker multi-stage build to minimize the size of the final image.
# This is the final stage that will produce the final Docker image.
FROM scratch

# Copy the statically-linked binary from the builder stage.
COPY --from=builder /app/revenueCalculator /revenueCalculator

# Command to run the executable
ENTRYPOINT ["/revenueCalculator"]
