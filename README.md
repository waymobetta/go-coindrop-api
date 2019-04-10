# go-coindrop-api

> The CoinDrop API

[![Go Report](https://goreportcard.com/badge/github.com/waymobetta/go-coindrop-api)](https://goreportcard.com/badge/github.com/waymobetta/go-coindrop-api)

## Getting started

Start server

```bash
make start
```

### Environments

export env vars in `.env.local`

example:

```
export POSTGRES_HOST="localhost"
```

then run:

```bash
make start/local
```

export env vars in `.env.staging`, then run:

```bash
make start/staging
```

export env vars in `.env.prod`, then run:

```bash
make start/prod
```

### Goa

generate controllers


```bash
make goa
```

## License

[-](LICENSE)
