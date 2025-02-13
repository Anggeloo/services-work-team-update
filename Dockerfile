FROM golang:1.23.5
WORKDIR /app/service-work-team-update
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 8182
CMD ["./main"]
