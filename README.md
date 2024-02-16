POC Valorant Client API Wrapper written in Go.

```go
    // Login to valorant account using client api
    user, err := govalo.Setup("na", USERNAME, PASSWORD)
    if (err != nil) {
        fmt.Println(err)
        return
    }

    // Prints user in format GAMENAME#TAGLINE
    fmt.Println(user)
```