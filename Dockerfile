# Stage 1: compile the program
FROM golang:1.17 as build-stage

ENV GIN_MODE=release
ENV PORT=1337

WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o server main.go

# Stage 2: build the image
FROM alpine:latest  
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /app/
COPY --from=build-stage /app/server .
EXPOSE 8080
CMD ["./server"]  