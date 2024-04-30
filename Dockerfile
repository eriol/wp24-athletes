FROM docker.io/golang:1.22.2-alpine3.19 AS builder

WORKDIR /app
ENV CGO_ENABLED=1

RUN apk -U --no-cache add git gcc musl-dev sqlite
COPY . .
RUN go build
RUN sqlite3 athletes.sqlite ".read extras/athletes.sql"


FROM docker.io/alpine:3.19

LABEL LastUpdate="2024-04-30"
COPY --from=builder /app/wp24-athletes /wp24-athletes
COPY --from=builder /app/images /images
COPY --from=builder /app/athletes.sqlite /athletes.sqlite
RUN apk -U --no-cache upgrade

EXPOSE 8080

ENTRYPOINT ["/wp24-athletes"]
