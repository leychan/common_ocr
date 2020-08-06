package baidu_ocr

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	file2 "github.com/leychan/common_ocr/cache/file"
	redis2 "github.com/leychan/common_ocr/cache/redis"
)

type accessToken struct {
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int64  `json:"expires_in"`
	SessionKey       string `json:"session_key"`
	AccessToken      string `json:"access_token"`
	Scope            string `json:"scope"`
	SessionSecret    string `json:"session_secret"`
	ErrorDescription string `json:"error_description"`
	Error            string `json:"error"`
}

var (
	goodsCachePath          string
	textCachePath           string
	goodsAccessTokenInRedis string
	textAccessTokenInRedis  string
	cacheDriver             string
	redisKeyMap             = map[string]string{}
	fileKeyMap              = map[string]string{}
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	goodsCachePath = os.Getenv("FILE_CACHE_PATH") + "/" + os.Getenv("FILE_GOODS_ACTOKEN_PATH")
	textCachePath = os.Getenv("FILE_CACHE_PATH") + "/" + os.Getenv("FILE_TEXT_ACTOKEN_PATH")
	goodsAccessTokenInRedis = os.Getenv("REDIS_GOODS_ACTOKEN_KEY")
	textAccessTokenInRedis = os.Getenv("REDIS_TEXT_ACTOKEN_KEY")
	cacheDriver = os.Getenv("CACHE_DRIVER")
	redisKeyMap = map[string]string{
		"goods": goodsAccessTokenInRedis,
		"text":  textAccessTokenInRedis,
	}
	fileKeyMap = map[string]string{
		"goods": goodsCachePath,
		"text":  textCachePath,
	}
}

//从百度api获取access_token
func GetAccessToken(t string) (string, error) {
	//获取缓存中的access_token等数据
	accessToken, err := getCachedAccessToken(t)
	if err == nil {
		return accessToken, nil
	}
	fmt.Println("取 http token")
	return getHttpAccessToken(t)
}

func getCachedAccessToken(t string) (string, error) {
	if cacheDriver == "redis" {
		if key, ok := redisKeyMap[t]; ok {
			token, err := getAccessTokenFromRedis(key)
			if err == nil {
				return token, nil
			}
		}
	} else {
		if key, ok := fileKeyMap[t]; ok {
			token, err := getAccessTokenFromFile(key)
			if err == nil {
				return token, nil
			}
		}
	}
	return "", errors.New("env params err")
}

func getHttpAccessToken(t string) (string, error) {
	b, err := getAccessApiUrl(t)
	if err != nil {
		return "", err
	}
	resp, err := http.Get(b)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	ac, err := unmarshalToken(body)
	if err != nil {
		return "", err
	}
	//正确返回
	if ac.Error == "" {
		if cacheDriver == "redis" {
			if key, ok := redisKeyMap[t]; ok {
				_ = storeAccessTokenToRedis([]byte(ac.AccessToken), key, ac.ExpiresIn)
			}
		} else {
			if key, ok := fileKeyMap[t]; ok {
				_ = storeAccessTokenToFile(body, key)
			}
		}
	} else {
		return "", errors.New(ac.Error)
	}
	return ac.AccessToken, nil
}

func getAccessTokenFromRedis(key string) (string, error) {
	return redis2.Get(key)
}

//获取本地缓存的token等相关数据
func getAccessTokenFromFile(cachePath string) (string, error) {
	c, err := file2.Get(cachePath)
	if err != nil {
		return "", err
	}
	ac, _ := unmarshalToken(c)
	file, _ := os.Open(cachePath)
	fileInfo, _ := file.Stat()
	//fmt.Println("now timestamp ", time.Now().Unix())
	//fmt.Println("file last modify time ", fileInfo.ModTime().Unix())
	if (time.Now().Unix() - fileInfo.ModTime().Unix()) > ac.ExpiresIn {
		err := os.Remove(cachePath)
		if err != nil {
			return "", err
		}
		return "", errors.New("token file is expired")
	}
	//fmt.Println("读取到本地缓存文件...")
	return ac.AccessToken, nil
}

func storeAccessTokenToRedis(body []byte, key string, expire int64) error {
	_, err := redis2.Set(key, body, expire)
	if err != nil {
		return err
	}
	return nil
}

func storeAccessTokenToFile(body []byte, path string) error {
	return file2.Set(path, body)
}

//解析返回的token等相关数据
func unmarshalToken(token []byte) (accessToken, error) {
	v := accessToken{}
	err := json.Unmarshal(token, &v)
	if err != nil {
		fmt.Println(err)
		return v, err
	}
	return v, nil
}
