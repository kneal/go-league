# go-league

[![Go Report Card](https://goreportcard.com/badge/github.com/kneal/go-league)](https://goreportcard.com/report/github.com/kneal/go-league)

Go client library for [League of Legends (LOL) API](https://developer.riotgames.com/apis)

## Documentation

### Usage

```go
import "github.com/kneal/go-league/league"
```

Construct a new League client, then use the various services on the client to access different parts of the LOL API. For example:

```go
client, _ := league.NewClient("league.url.com", nil)

// list champion rotations, including free-to-play and low-level free-to-play rotations
champions, _, err := client.Champion.GetAll()
```

_For more [examples]() see [Go doc](https://pkg.go.dev)_

### Authentication

The `league` package allows you to pass an [API Key](https://developer.riotgames.com/docs/portal#web-apis_api-keys) for authenticating to LOL.

Example using API Key:

```go
client, _ := league.NewClient("league.url.com", nil)

client.Authentication.SetTokenAuth("token")
```


## Contributing

Always welcome new PRs! See [_Contributing_](.github/CONTRIBUTING.md) for further instructions.

## Bugs and Feature Requests

Found something that doesn't seem right or have a feature request? [Please open a new issue](../../issues/new/).
