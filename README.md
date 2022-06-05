[![Maintainability](https://api.codeclimate.com/v1/badges/f2b7b99e9c1d1c93ca47/maintainability)](https://codeclimate.com/github/nvzard/Casino-Royale/maintainability)

<p align="center">
<img width="668" alt="Screenshot 2022-05-17 at 1 08 21 PM" src="https://user-images.githubusercontent.com/19310512/168756153-2892a485-d826-4e5d-bfc6-9e97c8110f31.png">
</p>

<!-- An API to simulate a deck of cards -->
## Documentation

- [Summary](#summary)
- [Development Setup](#development)
- [API Documentation](#api-documentation)
- [Future Improvements](#api-documentation)
- [Side Note](#side-note)

---
## Summary

This project exposes an API to simulate deck of cards which can be used for card games.

## Tech Stack

Written in Go (Golang 1.18) using [gin](https://github.com/gin-gonic/gin) as web framework,[Postgres](https://github.com/postgres/postgres) as database, [zap](https://github.com/uber-go/zap) for logging, [prometheus](https://github.com/prometheus/prometheus) and [grafana](https://github.com/grafana/grafana) for monitoring.

## Development

### Requirements

```
git
docker
docker-compose
docker-sync (0.5.14) [Use `gem install docker-sync -v 0.5.14` for mac]
```

### Getting Started with Development

1. Clone repository

```
git clone https://github.com/nvzard/Casino-Royale.git
```

2. Change directory to the cloned repository

```
cd Casino-Royale
```

3. Build image

```
make build
```

4. Run start command

```
make start
```

### Run Tests

```
make test
```

### Remove all artifacts and dependencies

```
make clean
```
---

## API Documentation

```
GET   /healthcheck          # ok
POST  /deck                 # create new deck (query parameters: shuffle, cards)
GET   /deck/:deck_id        # open deck by id
POST  /deck/:deck_id/draw   # draw cards from the deck (query parameters: count)
```

### 1. Create a Deck

```
POST /deck
```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `shuffle` | `bool`   | **Optional**. true or false |
| `cards`   | `string` | **Optional**. Custom Cards Code: AS,1S,2S, etc |


#### Example Request 1
```
POST /deck
```

#### Response
```
{
    "deck_id": "a251071b-662f-44b6-ba11-e24863039c59",
    "shuffled": false,
    "remaining": 52
}
```

#### Example Request 2

```
POST /deck?shuffle=true&cards=AS,2S
```

#### Response
```
{
    "deck_id": "9bc9d90e-3056-40c7-94bc-241915cbee4d",
    "shuffled": true,
    "remaining": 2
}
```

### 2. Open a Deck

```
GET /deck/:deck_id
```

#### Example Request 1

```
GET /deck/a251071b-662f-44b6-ba11-e24863039c59
```

#### Response
```
{
    "deck_id": "a251071b-662f-44b6-ba11-e24863039c59",
    "shuffled": false,
    "remaining": 3,
    "cards": [
        {
            "value": "ACE",
            "suit": "SPADES",
            "code": "AS"
        },
	    {
            "value": "KING",
            "suit": "HEARTS",
            "code": "KH"
        },
        {
            "value": "8",
            "suit": "CLUBS",
            "code": "8C"
        }
    ]
}
```

### 3. Draw a Card

```
POST /deck/:deck_id/draw
```


| Parameter | Type | Description |
| :--- | :--- | :--- |
| `count` | `string`   | **Optional**. Number of cards to draw [default=1]|


#### Example Request 1

```
POST /deck/a251071b-662f-44b6-ba11-e24863039c59/draw
```

#### Response
```
{
    "cards": [
        {
            "value": "QUEEN",
            "suit": "HEARTS",
            "code": "QH"
        }
    ]
}
```
#### Example Request 2

```
POST /deck/:deck_id/draw?count=2
```

#### Response
```
{
    "cards": [
	    {
            "value": "KING",
            "suit": "HEARTS",
            "code": "KH"
        },
        {
            "value": "8",
            "suit": "CLUBS",
            "code": "8C"
        }
    ]
}
```

---
## Future Improvements

Some of the improvements that can be made but were skipped due to time constraints.

- Write more tests and improve code coverage.
- Improve controller function logic
- Add cleaner error handling mechanism and single data store for all error messages.
- Setup CI/CD pipeline for testing and deployment.
- Add some kind of authentication mechanism.

---
## Side Note

This was my first ever go project. Please open an issue and let me know incase I messed something up. Some feedback would really help me up my go game.
