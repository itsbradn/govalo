<p align="center">
	<h1 align="center">GoValo</h1>
	<p align="center">
		<b>GoValo</b> is a Valorant Client API library built for go. Designed for simplicity and ease of use you can get up in running in a few lines of code.
	</p>
</p>

<p align="center">
	<a aria-label="Package License" href="https://pkg.go.dev/github.com/itsbradn/govalo">
		<img alt="" src="https://img.shields.io/github/license/itsbradn/govalo?style=for-the-badge&labelColor=000000">
	</a>
	<a aria-label="Package Version" href="https://pkg.go.dev/github.com/itsbradn/govalo">
		<img alt="" src="https://img.shields.io/github/v/release/itsbradn/govalo?style=for-the-badge&labelColor=000000">
	</a>
	<a aria-label="Valorant API Coverage" href="https://pkg.go.dev/github.com/itsbradn/govalo">
		<img alt="" src="https://img.shields.io/badge/COVERAGE-13%2F31%20Endpoints-purple?style=for-the-badge&labelColor=%23000000">
	</a>
</p>

## ⚡️ Get Started

```go
package main

import (
	"fmt"

	"github.com/itsbradn/govalo"
)

func main() {
	api, err := govalo.Setup("na", USERNAME, PASSWORD)

	if err != nil {
		fmt.Print(err)
		return
	}

	userinfo, err := valapi.GetUserInfo()

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%s#%s\n", userinfo.Account.GameName, userinfo.Account.TagLine)
}
```

## 🎯 Endpoints

### Match History

```go
func main() {
	api, err := govalo.Setup("na", USERNAME, PASSWORD)
	if err != nil {
		fmt.Print(err)
		return
	}

	userinfo, err := valapi.GetUserInfo()
	if err != nil {
		fmt.Print(err)
		return
	}

	history, err := valapi.GetMatchHistory(userinfo.PlayerUUID, &govalo.MatchHistoryOptions{
		StartIndex: 0
		EndIndex:   5
	})
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("Most recent match: %s\n", history.History[0].MatchID)
}
```

### Competitive Updates

```go
func main() {
	api, err := govalo.Setup("na", USERNAME, PASSWORD)
	if err != nil {
		fmt.Print(err)
		return
	}

	userinfo, err := valapi.GetUserInfo()
	if err != nil {
		fmt.Print(err)
		return
	}

	competitiveUpdates, err := valapi.GetCompetitiveUpdates(userinfo.PlayerUUID, &govalo.CompetitiveUpdatesOptions{
		StartIndex: 0
		EndIndex:   5
	})
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("Most recent competitive update: %s\n", competitiveUpdates.Matches[0])
}
```
