package baidu_ocr

import (
	"errors"
)

var requestUrl = map[string]string{
	"general_basic":    "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic",    //通用版
	"accurate_basic":   "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic",   //通用高精度版,
	"idcard":           "https://aip.baidubce.com/rest/2.0/ocr/v1/idcard",           //身份证
	"bankcard":         "https://aip.baidubce.com/rest/2.0/ocr/v1/bankcard",         //银行卡
	"business_license": "https://aip.baidubce.com/rest/2.0/ocr/v1/business_license", //营业执照
	"business_card":    "https://aip.baidubce.com/rest/2.0/ocr/v1/business_card",    //名片
	"receipt":          "https://aip.baidubce.com/rest/2.0/ocr/v1/receipt",          //通用票据
	"handwriting":      "https://aip.baidubce.com/rest/2.0/ocr/v1/handwriting",      //手写文字
	"formula":          "https://aip.baidubce.com/rest/2.0/ocr/v1/formula",          //公式,
}

var TextOcrType = map[string]string{
	"general_basic":    "通用版",
	"accurate_basic":   "通用高精度版", //通用高精度版,
	"idcard":           "身份证",    //身份证
	"bankcard":         "银行卡",    //银行卡
	"business_license": "营业执照",   //营业执照
	"business_card":    "名片",     //名片
	"receipt":          "通用票据",   //通用票据
	"handwriting":      "手写文字",   //手写文字
	"formula":          "公式",     //公式,
}

var TextTypeSlice = []string{"general_basic", "accurate_basic", "idcard", "bankcard", "business_license",
	"business_card", "receipt", "handwriting", "formula"}

func DoTextOcr(base64str string, t string) ([]byte, error) {
	url, ok := requestUrl[t]
	if !ok {
		return nil, errors.New("unknown identity type")
	}
	token, _ := GetAccessToken("text")
	url += "?access_token=" + token + "&baike_num=3"
	params := []byte("image=" + base64str)
	b, err := formRequest(params, url)
	if err != nil {
		return nil, err
	}
	return b, nil
}
