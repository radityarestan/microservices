FROM golang:1.18-alpine
RUN mkdir /app
WORKDIR /app
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 go build -o chat-service ./cmd

CMD [ "./chat-service" ]