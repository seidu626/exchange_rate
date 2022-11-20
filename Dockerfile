FROM golang:1.19-buster as builder

# Create and change to the app directory.
WORKDIR /app

COPY .env config.yaml app_server ./

# Copy local code to the container image.
COPY . ./

RUN mv /app/app_server /usr/bin/

ENTRYPOINT ["/usr/bin/app_server"]