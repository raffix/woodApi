# Use an existing Go image as a base
FROM golang:1.19.0

# Set the working directory in the container
WORKDIR /app

# Copy the application code to the container
COPY . .

# Build the application
RUN go build -mod=vendor -o main .

# Expose port 8080 for the web API
EXPOSE 8080

# Start the web API
CMD ["./main"]