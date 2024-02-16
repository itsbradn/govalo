package conf

import (
	"encoding/base64"
	"encoding/json"
)

type ClientPlatform struct {
	Type    string `json:"platformType"`
	OS      string `json:"platformOS"`
	Version string `json:"platformOSVersion"`
	Chipset string `json:"platformChipset"`
}

var clientPlatformData = ClientPlatform{
	Type:    "PC",
	OS:      "Windows",
	Version: "10.0.19043.1.256.64bit",
	Chipset: "Unknown",
}

func GetClientPlatformEncoded() (string, error) {
	body, err := json.Marshal(clientPlatformData)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(body), nil
}

func GetClientPlatform() *ClientPlatform {
	return &clientPlatformData
}

func SetClientPlatform(platform ClientPlatform) *ClientPlatform {
	clientPlatformData = platform
	return &clientPlatformData
}
