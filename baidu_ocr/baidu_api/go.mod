module github.com/leychan/common_ocr/baidu_ocr/baidu_api

go 1.14

require (
	github.com/leychan/common_ocr/baidu_ocr/baidu_config v0.0.0
	github.com/leychan/common_ocr/cache/redis v0.0.0
)

replace (
	github.com/leychan/common_ocr/baidu_ocr/baidu_config => ../baidu_config
	github.com/leychan/common_ocr/cache/redis => ../../cache/redis
)
