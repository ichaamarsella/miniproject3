# Use the official Golang image as the base
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port on which the app listens
EXPOSE 8080

# Define the command to run the executable
CMD ["./main"]
