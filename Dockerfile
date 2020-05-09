# Building enviroment
FROM golang:alpine as builder
LABEL maintainer="Alexander Zorkin"
# Git required for fetch dependenciesd
RUN apk update && apk add --no-cache git gcc libc-dev
WORKDIR /app
# Go mod required for faster build your application
COPY go.mod go.sum ./
RUN go mod download
# Start build application
COPY . .
# TODO: Modify build command for your application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Running enviroment
FROM alpine:latest
# For debugging
RUN apk --no-cache add bash
WORKDIR /app
# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main /app/.env ./

# Add block below if you run app on not Unix-like OS
# RUN apk --no-cache add dos2unix
# RUN dos2unix ./wait-for-it.sh

# Finish block
# Specify port of your application
EXPOSE 6060
CMD ./main