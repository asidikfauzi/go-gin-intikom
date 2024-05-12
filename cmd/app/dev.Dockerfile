FROM golang:1.22.2-alpine
WORKDIR /app
COPY . ./

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]