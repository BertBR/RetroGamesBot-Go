# RetroGamesBot-Go (WIP)

- A rewrite of [RetroGamesBot](https://github.com/BertBR/RetroGamesBot) in Golang.

### Setup

> For `/top10` command, you must have same database schema like [Game Struct](https://github.com/BertBR/RetroGamesBot-Go/blob/main/pkg/database/main.go#L11-L19).

Ensure you have a `.env` file like `.env.example` with your `BOT_TOKEN`, `WEBHOOK_URL` and `PORT`

| ENV   |      What  :|
|----------|:-------------:|
| BOT_TOKEN |  Get your bot token in [Bot Father](https://t.me/BotFather) |
| WEBHOOK_URL |    Any test url you want (ngrok, localtunnel, etc)   |
| PORT | default: 3000 |

---
### How to Run
`make dev`