# Step 1: Use Go 1.23 base image
FROM golang:1.23-bullseye AS build

# Step 2: Install necessary packages
RUN apt-get update && apt-get install -y bash git curl && rm -rf /var/lib/apt/lists/*

# Step 3: Install Air using the correct module path
RUN go install github.com/air-verse/air@latest

# Step 4: Set the working directory
WORKDIR /app

# Step 5: Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Step 6: Copy the application source code
COPY . .

# Step 7: Expose the application port
EXPOSE 8080

# Step 8: Command to run Air for live reloading
CMD ["air"]
