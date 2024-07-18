package crawlerxueqiu

import (
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36"
)

func GetHttpClient() *resty.Client {
	client := resty.New().
		SetRetryCount(3).
		SetRetryWaitTime(1*time.Second).
		SetRetryMaxWaitTime(10*time.Second).
		SetHeader("User-Agent", userAgent).
		SetCookieJar(nil)

	return client
}
