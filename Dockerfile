FROM golang:1.22.3-alpine3.19

RUN addgroup app && adduser -S -G app app

USER app
WORKDIR /app

COPY go.mod go.sum ./

USER root
RUN chown -R app:app /app
RUN chmod -R 775 /app
USER app

RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]