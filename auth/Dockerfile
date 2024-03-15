# base image
FROM golang:latest AS builder
# create appuser
RUN adduser --disabled-login appuser
# create workspace
WORKDIR /app
COPY go.mod go.sum ./
# fetch dependancies
RUN go mod download && \
    go mod verify
# copy the source code as the last step
COPY . .
# build binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/course-shop-auth ./src


# build a small image
FROM alpine:3.17.3
# import the user and group files from the builder
COPY --from=builder /etc/passwd /etc/passwd
# copy the static executable
COPY --from=builder --chown=appuser:1000 /go/bin/course-shop-auth /course-shop-auth
# use a non-root user
USER appuser
# expose port
EXPOSE 3000
# run app
ENTRYPOINT ["./course-shop-auth"]