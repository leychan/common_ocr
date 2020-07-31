package baidu_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/leychan/common_ocr/baidu_ocr/baidu_config"
	file2 "github.com/leychan/common_ocr/cache/file"
	redis2 "github.com/leychan/common_ocr/cache/redis"
)

//var goodsCachePath = "goods_cache_file"
//var textCachePath = "text_cache_file"
var goodsAccessTokenInRedis = "goods_access_token"
var textAccessTokenInRedis = "text_access_token"

//从百度api获取access_token
func GetAccessToken(t string) (string, error) {
	fmt.Println("t:", t)
	//获取缓存中的access_token等数据
	switch strings.ToLower(t) {
	case "goods":
		token, err := getAccessTokenFromRedis(goodsAccessTokenInRedis)
		if err == nil {
			return token, nil
		}
	case "text":
		token, err := getAccessTokenFromRedis(textAccessTokenInRedis)
		if err == nil {
			return token, nil
		}
	default:
		return "", errors.New("unknown env params")
	}
	fmt.Println("取 http token")
	return getHttpAccessToken(t)
}

func getHttpAccessToken(t string) (string, error) {
	b, err := baidu_config.GetInstance(t)
	fmt.Println(err)
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
		fmt.Println(err)
		return "", err
	}
	ac, err := unmarshalToken(body)
	if err != nil {
		return "", err
	}
	//正确返回
	if ac.Error == "" {
		switch strings.ToLower(t) {
		case "goods":
			_, err = storeAccessTokenToRedis([]byte(ac.AccessToken), goodsAccessTokenInRedis, ac.ExpiresIn)
		case "text":
			_, err = storeAccessTokenToRedis([]byte(ac.AccessToken), textAccessTokenInRedis, ac.ExpiresIn)
		}
	} else {
		return "", errors.New(ac.Error)
	}
	return ac.AccessToken, nil
}

func getAccessTokenFromRedis(key string) (string, error) {
	return redis2.Get(key)
}

//func getCachedTextAccessToken() (string, error) {
//	return getCachedAccessToken(textCachePath)
//}
//
//func getCachedGoodsAccessToken() (string, error) {
//	return getCachedAccessToken(goodsCachePath)
//}

//获取本地缓存的token等相关数据
func getFileCachedAccessToken(cachePath string) (string, error) {
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

func storeAccessTokenToRedis(body []byte, key string, expire int64) (bool, error) {
	_, err := redis2.Set(key, body, expire)
	if err != nil {
		return false, err
	}
	return true, err
}

//func storeAccessTokenToFile(path string, body []byte) (bool, error) {
//	return file2.Set(path, body)
//}

//解析返回的token等相关数据
func unmarshalToken(token []byte) (baidu_config.AccessToken, error) {
	v := baidu_config.AccessToken{}
	err := json.Unmarshal(token, &v)
	if err != nil {
		fmt.Println(err)
		return v, err
	}
	return v, nil
}
