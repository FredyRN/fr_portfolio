FROM golang:1.19-alpine3.16 AS builder

LABEL maintainer = "Fredy Rodriguez - software@jfredy-rodriguez.com"

# Install container dependencies
RUN apk update \
    && apk add --no-cache git \
    && apk add --no-cache bash \
    && apk add --no-cache ca-certificates \
    && apk add --update gcc musl-dev \
    && apk add build-base \
    && update-ca-certificates

# Creates and select workdir folder
WORKDIR /go/src/fr_projects

# Download dependencies
COPY backend/go.mod ./
RUN go mod download && go mod verify

# Copy the source from current folder to workdir inside the container
COPY . .
COPY .env .

# Build project
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -installsuffix cgo -o fr_portfolio .

# Use linux alpine
FROM alpine:3.16

RUN apk update \
    && apk add --no-cache git \
    && apk add --no-cache bash \
    && apk add --no-cache ca-certificates \
    && apk add --update gcc musl-dev \
    && apk build-base \
    && update-ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/fr_projects/fr_portfolio .

EXPOSE 8000

CMD [ "./fr_portfolio" ]
