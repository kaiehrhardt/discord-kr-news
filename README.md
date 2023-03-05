# discord kartrider news publisher bot

Simple discord bot to post kartider news in a selected channel.

## config via env vars

|env var|description|default|
|--|--|--|
| KRN_DISCORD_TOKEN | discord auth token, [docs](https://discord.com/developers/docs/topics/oauth2) | `dummy` |
| KRN_DISCORD_CHANNEL | discord channel id, [docs](https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID-) | `dummy` |
| KRN_CHECK_INTERVAL | duration between checks, [docs](https://pkg.go.dev/time#example-ParseDuration) | `600s` |

## container run

```bash
docker pull ghcr.io/kaiehrhardt/discord-kr-news:latest
docker run --rm \
  --env KRN_DISCORD_TOKEN=<token> \
  --env KRN_DISCORD_CHANNEL=<channelid> \
  ghcr.io/kaiehrhardt/discord-kr-news:latest
```

## helm

wip
