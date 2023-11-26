FROM golang:latest
LABEL maintainer="Kristian Poslek <kristian@poslek.com>"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
COPY . .

EXPOSE 8080
CMD ["air"]