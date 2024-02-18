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

// Fetch previous 10 matches
history, err := valapi.GetMatchHistory("0e5b44a6-3085-545e-afb3-6c296e2c494e", nil)
if err != nil {
	fmt.Print(err)
	return
}

firstMatch := history.History[0] // Get the most recent match

matchDetails, err := valapi.GetMatchDetails(firstMatch.MatchID)
if err != nil {
	fmt.Print(err)
	return
}

fmt.Print(matchDetails.MatchInfo)
```
