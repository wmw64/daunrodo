FROM golang:alpine AS build

RUN apk add --update --no-cache tzdata

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" ./cmd/daunrodo

FROM alpine
#FROM scratch

ENV TZ Europe/Moscow

#COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
#COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#COPY --from=build /etc/passwd /etc/passwd
#COPY --from=build /etc/group /etc/group

WORKDIR /app
COPY --from=build /app/daunrodo /app/daunrodo
COPY --from=build /app/entrypoint.sh /app/entrypoint.sh
COPY --from=build /app/config/config.yaml /app/config.yaml

USER 1000:1000

ENTRYPOINT ["sh", "entrypoint.sh"]
