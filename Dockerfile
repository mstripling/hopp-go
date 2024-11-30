
# Use the Go image as the base
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the entire application source code
COPY . .

# Install gettext package for envsubst
RUN apt-get update && apt-get install -y gettext

# Make the script executable
RUN chmod +x generate-init.sh

# Run the script to generate the `init.sql` file
RUN ./generate-init.sh

# Set the environment variable for the application port
ENV PORT=80

# Expose the port
EXPOSE 80

# Command to run the application
ENTRYPOINT ["/app/main"]

