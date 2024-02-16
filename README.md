POC Valorant Client API Wrapper written in Go.

```go
    // Login to valorant account using client api
    valapi, err := govalo.Setup("na", USERNAME, PASSWORD)
    if (err != nil) {
        fmt.Println(err)
        return
    }

	userinfo, err := valapi.GetUserInfo()
	if err != nil {
		fmt.Print(err)
		return
	}

    // Prints user in format GAMENAME#TAGLINE
	fmt.Printf("%s#%s\n", userinfo.Account.GameName, userinfo.Account.TagLine)
```