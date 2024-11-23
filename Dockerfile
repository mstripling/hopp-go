
# Use the Go image as the base
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the entire application source code
COPY . .

# Set the environment variable for the application port
ENV PORT=80

# Expose the port
EXPOSE 80

# Command to run the application
ENTRYPOINT ["/app/main"]

