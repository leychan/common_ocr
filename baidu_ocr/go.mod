module github.com/leychan/common_ocr/baidu_ocr/baidu_api

go 1.14

require (
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/leychan/common_ocr/cache/file v0.0.0-20200731031825-6aa139509aa5 // indirect
	github.com/leychan/common_ocr/cache/redis v0.0.0-20200731030744-e433e53f2a30
	github.com/leychan/go-helper v0.0.0-20200803013805-6c3027c573d3 // indirect
)

replace github.com/leychan/common_ocr/cache/file => ./../cache/file
