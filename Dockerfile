#############################################
# STEP 1 build executable binary & migrate db
#############################################
FROM adhiva/alpine-golang-migrate as builder

# Set your working directory, as your project folder
WORKDIR /app/go-jlpt-n5

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build Go Binary
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o main .

# ############################
# # STEP 2 build a small image
# ############################
FROM alpine:latest

# Set Timezone
# ENV TIMEZONE Asia/Jakarta
ENV GIN_MODE release

COPY config.yaml /
COPY --from=builder /app/go-jlpt-n5/main ./main

# Expose port 8080 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]