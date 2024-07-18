package crawlerxueqiu

import (
	"testing"
)

func TestCrawler_GetStockBasicData(t *testing.T) {
	c := NewCrawler()

	data, err := c.GetStockBasicData("腾讯")
	if err != nil {
		t.Error(err)
	}
	t.Log(data.Format())

	data, err = c.GetStockBasicData("gzmt")
	if err != nil {
		t.Error(err)
	}
	t.Log(data.Format())

	data, err = c.GetStockBasicData("coinbase")
	if err != nil {
		t.Error(err)
	}
	t.Log(data.Format())
}
