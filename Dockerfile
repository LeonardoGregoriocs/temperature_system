FROM golang:latest

WORKDIR /app

RUN go mod tidy

COPY . .

EXPOSE 8080

COPY entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["entrypoint.sh"]