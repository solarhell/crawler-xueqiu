package main

import (
	"fmt"
	"log"
	"net/http"
)

func getToken() (string, error) {
	resp, err := http.Get("https://xueqiu.com/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	for _, cookie := range resp.Cookies() {
		fmt.Println(cookie.Name, cookie.Value)
		if cookie.Name == "xq_a_token" {
			return cookie.Value, nil
		}
	}

	return "", nil
}

func getSuggestStock(q string) (string, error) {
	token, err := getToken()
	if err != nil {
		return "", err
	}

	url := "https://xueqiu.com/query/v1/suggest_stock.json?q=" + q
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Cookie", token)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		// 处理响应
		// 假设响应是 JSON 格式，解析 JSON 并返回数据
		// 这里为了简单起见，仅返回空字符串
		return "", nil
	} else {
		return "", fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}
}

func getStockBasicData(symbol string) error {
	code, err := getSuggestStock(symbol)
	if err != nil {
		return err
	}

	if code == "" {
		return fmt.Errorf("没有返回建议股票代码")
	}

	// 使用获取到的代码进行其他操作
	// 比如获取股票详细数据
	// 这里为了简单起见，仅打印代码
	fmt.Println("Suggest Stock Code:", code)

	return nil
}

func main() {
	symbol := "gzmt"
	err := getStockBasicData(symbol)
	if err != nil {
		log.Fatalln(err)
	}
}
