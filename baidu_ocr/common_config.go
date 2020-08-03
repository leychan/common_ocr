package baidu_ocr

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type AccessToken struct {
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int64  `json:"expires_in"`
	SessionKey       string `json:"session_key"`
	AccessToken      string `json:"access_token"`
	Scope            string `json:"scope"`
	SessionSecret    string `json:"session_secret"`
	ErrorDescription string `json:"error_description"`
	Error            string `json:"error"`
}

type account struct {
	apiKey    string
	apiSecret string
}

func getAccessApiUrl(t string) (string, error) {
	a, err := getBaiduOCRAccount(t)
	if err != nil {
		return "", err
	}
	accessUrl := "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=" +
		a.apiKey + "&client_secret=" + a.apiSecret
	return accessUrl, nil
}

func getBaiduOCRAccount(t string) (account, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return account{}, err
	}
	a := account{}
	switch strings.ToLower(t) {
	case "goods":
		a.apiKey = os.Getenv("BAIDU_GOODS_APIKEY")
		a.apiSecret = os.Getenv("BAIDU_GOODS_APISECRET")
	case "text":
		a.apiKey = os.Getenv("BAIDU_TEXT_APIKEY")
		a.apiSecret = os.Getenv("BAIDU_TEXT_APISECRET")
	default:
		return account{}, errors.New("unknown env params")
	}
	return a, nil
}
