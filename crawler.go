package crawlerxueqiu

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Crawler struct {
	httpClient *resty.Client

	cookies []*http.Cookie
}

type Interface interface {
	GetCookie() error
	GetSuggestStock(query string) (*SuggestStockData, error)
	GetStockBasicData(symbol string) (*QuoteData, error)
}

func NewCrawler() *Crawler {
	return &Crawler{
		httpClient: GetHttpClient(),
	}
}

func (i *Crawler) GetCookies() error {
	resp, err := i.httpClient.R().Get("https://xueqiu.com/")
	if err != nil {
		return err
	}

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "xq_a_token" && cookie.Value != "" {
			cookies := make([]*http.Cookie, 0)
			cookies = append(cookies, cookie)
			i.cookies = cookies
			return nil
		}
	}

	return fmt.Errorf("获取cookie失败")
}

func (i *Crawler) GetSuggestStock(q string) (*SuggestStockData, error) {
	err := i.GetCookies()
	if err != nil {
		return nil, err
	}

	suggestStock := new(SuggestStock)

	resp, err := i.httpClient.R().
		SetCookies(i.cookies).
		SetResult(suggestStock).
		SetQueryParam("q", q).
		Get("https://xueqiu.com/query/v1/suggest_stock.json")
	if err != nil {
		return nil, err
	}

	if suggestStock.Code != 200 || !suggestStock.Success {
		return nil, fmt.Errorf("获取建议股票失败: %s", resp.String())
	}

	if len(suggestStock.Data) == 0 {
		return nil, fmt.Errorf("没有找到建议股票: %s", resp.String())
	}

	return &suggestStock.Data[0], nil
}

func (i *Crawler) GetStockBasicData(symbol string) (*QuoteData, error) {
	code, err := i.GetSuggestStock(symbol)
	if err != nil {
		return nil, err
	}

	quote := new(Quote)

	resp, err := i.httpClient.R().
		SetCookies(i.cookies).
		SetResult(quote).
		SetQueryParams(map[string]string{
			"symbol": code.Code,
			"extend": "detail",
		}).
		Get("https://stock.xueqiu.com/v5/stock/quote.json")
	if err != nil {
		return nil, err
	}

	if quote.ErrorCode != 0 {
		return nil, fmt.Errorf("获取quote失败: %s", resp.String())
	}

	return &quote.Data, nil
}
