FROM golang:latest

WORKDIR /app

COPY . .

EXPOSE 8000

COPY entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["entrypoint.sh"]