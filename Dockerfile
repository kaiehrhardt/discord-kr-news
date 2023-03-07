FROM golang:1.19-alpine as build
WORKDIR /app
COPY . .
RUN go mod download && CGO_ENABLED=0 go build

FROM gcr.io/distroless/static-debian11
COPY --from=build /app/discord-kr-news /
ENV KRN_DISCORD_TOKEN=dummy
ENV KRN_DISCORD_CHANNEL=dummy
ENV KRN_CHECK_INTERVAL=600s
ENTRYPOINT ["discord-kr-news"]
