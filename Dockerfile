# Base Image
FROM golang:1.12.0-alpine3.9

# We Create an /app directory within our
# Image that will hold our application source
WORKDIR /app

# We Copy everything in the root directory
# into our WORKDIR
COPY . .

# Run go build to compile the binary
# execute of our go program
RUN go build -o main .

# Our start command which kick off
# Create binary executable
CMD ["/app/main"]


