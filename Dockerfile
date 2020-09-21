FROM golang:1.14.2-alpine3.11 as builder


RUN mkdir -p /go/src/echo-jwt/ && mkdir -p /log && mkdir -p /root/.ssh

RUN apk --no-cache add \ 
		git \
		openssh \
		openssh-server

# Set the Current Working Directory inside the container
WORKDIR /go/src/echo-jwt/

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

#Build dependenices
RUN sh packages.sh

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

##### new stage to copy the artifact #####

FROM alpine:3.11

RUN mkdir -p /echo && mkdir -p /log

# Set the Current Working Directory inside the container
WORKDIR /echo

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/src/echo-jwt/main .

CMD ["./main"]