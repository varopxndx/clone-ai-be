FROM golang:1.18.3-alpine AS builder
WORKDIR /build

# Add gcc and libc-dev early so it is cached
RUN set -xe \
	&& apk add --no-cache gcc libc-dev

RUN apk update && apk add git

ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux

# Install dependencies earlier so they are cached between builds
COPY go.mod ./
RUN set -xe \
	&& go mod download

# Copy the source code, because directories are special, there are separate layers
COPY . .

# Get the version name and git commit as a build argument
ARG GIT_VERSION
ARG GIT_COMMIT

RUN [ "/bin/ls", "-ltra"]
# Build the application
RUN go build -trimpath -o clone-ai-be
# Let's create a /dist folder containing just the files necessary for runtime.
# Later, it will be copied as the / (root) of the output image.
WORKDIR /dist
RUN cp /build/clone-ai-be ./clone-ai-be
RUN cp /build/assets/chats.json ./chats.json

# Copy or create other directories/files your app needs during runtime.
# E.g. this example uses /data as a working directory that would probably
#      be bound to a perstistent dir when running the container normally

# start a new stage that copies in the binary built in the previous stage
# Create the minimal runtime image
FROM alpine:latest AS deploy

COPY --chown=0:0 --from=builder /dist /

EXPOSE 8080

WORKDIR /data

ENTRYPOINT ["/clone-ai-be"]
