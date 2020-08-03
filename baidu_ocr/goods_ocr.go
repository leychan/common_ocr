package baidu_ocr

import (
	"errors"
)

type GoodsParams struct {
	Image  string   `json:"image"`
	Scenes []string `json:"scenes"`
}

var requestUrl2 = map[string]string{
	"advanced_general":    "https://aip.baidubce.com/rest/2.0/image-classify/v2/advanced_general",
	"object_detect":       "https://aip.baidubce.com/rest/2.0/image-classify/v1/object_detect",
	"multi_object_detect": "https://aip.baidubce.com/rest/2.0/image-classify/v1/multi_object_detect",
	"animal":              "https://aip.baidubce.com/rest/2.0/image-classify/v1/animal",
	"plant":               "https://aip.baidubce.com/rest/2.0/image-classify/v1/plant",
	"logo_search":         "https://aip.baidubce.com/rest/2.0/image-classify/v2/logo",
	"ingredient":          "https://aip.baidubce.com/rest/2.0/image-classify/v1/classify/ingredient",
	"dishs":               "https://aip.baidubce.com/rest/2.0/image-classify/v2/dish",
	"red_wine":            "https://aip.baidubce.com/rest/2.0/image-classify/v1/redwine",
	"currency":            "https://aip.baidubce.com/rest/2.0/image-classify/v1/currency",
	"landmark":            "https://aip.baidubce.com/rest/2.0/image-classify/v1/landmark",
}

var GoodsOcrType = map[string]string{
	"advanced_general":    "通用版",
	"object_detect":       "图像单主体检测",
	"multi_object_detect": "图像多主体检测",
	"animal":              "动物",
	"plant":               "植物",
	"logo_search":         "logo",
	"ingredient":          "果蔬",
	"dishs":               "菜品",
	"red_wine":            "红酒",
	"currency":            "货币",
	"landmark":            "地标",
}

var GoodsTypeSlice = []string{"advanced_general", "plant", "animal", "dishs", "logo_search", "object_detect",
	"multi_object_detect", "ingredient", "red_wine", "currency", "landmark"}

func DoGoodsOcr(base64str string, t string) ([]byte, error) {
	url, ok := requestUrl2[t]
	if !ok {
		return nil, errors.New("unknown identity type")
	}
	token, _ := GetAccessToken("goods")
	url += "?access_token=" + token + "&baike_num=3"
	params := []byte("image=" + base64str)
	b, err := formRequest(params, url)
	if err != nil {
		return nil, err
	}
	return b, nil
}
