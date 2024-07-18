package crawlerxueqiu

import (
	"fmt"
	"strings"
)

type SuggestStock struct {
	Code    int                `json:"code"`
	Data    []SuggestStockData `json:"data"`
	Message string             `json:"message"`
	Success bool               `json:"success"`
}

type SuggestStockData struct {
	Code        string `json:"code"`
	Label       string `json:"label"`
	Query       string `json:"query"`
	State       int    `json:"state,omitempty"`
	StockType   int    `json:"stock_type,omitempty"`
	Type        int    `json:"type"`
	StatockType int    `json:"statock_type,omitempty"`
}

type Quote struct {
	Data             QuoteData `json:"data"`
	ErrorCode        int       `json:"error_code"`
	ErrorDescription string    `json:"error_description"`
}

type QuoteData struct {
	Market QuoteMarket `json:"market"`
	Quote  QuoteQuote  `json:"quote"`
	Others QuoteOthers `json:"others"`
	Tags   []QuoteTags `json:"tags"`
}

func (q *QuoteData) Format() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s(%s): %v (%v%%)\n", q.Quote.Name, q.Quote.Symbol, q.Quote.Current, q.Quote.Percent))
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("最高: %v\n", q.Quote.High))
	sb.WriteString(fmt.Sprintf("最低: %v\n", q.Quote.Low))
	sb.WriteString(fmt.Sprintf("总市值: %v\n", q.Quote.MarketCapital))
	sb.WriteString(fmt.Sprintf("市盈率(TTM): %v\n", q.Quote.PeTtm))
	sb.WriteString(fmt.Sprintf("股息率(TTM): %v\n", q.Quote.DividendYield))
	return sb.String()
}

type QuoteQuote struct {
	CurrentExt               float64 `json:"current_ext"`
	Symbol                   string  `json:"symbol"`
	PercentExt               float64 `json:"percent_ext"`
	Delayed                  int     `json:"delayed"`
	Type                     int     `json:"type"`
	TickSize                 float64 `json:"tick_size"`
	FloatShares              int64   `json:"float_shares"`
	High                     float64 `json:"high"`
	FloatMarketCapital       float64 `json:"float_market_capital"`
	LotSize                  int     `json:"lot_size"`
	LockSet                  int     `json:"lock_set"`
	Chg                      float64 `json:"chg"`
	Eps                      float64 `json:"eps"`
	LastClose                float64 `json:"last_close"`
	ProfitFour               float64 `json:"profit_four"`
	Volume                   int     `json:"volume"`
	VolumeRatio              float64 `json:"volume_ratio"`
	ProfitForecast           float64 `json:"profit_forecast"`
	TurnoverRate             float64 `json:"turnover_rate"`
	Name                     string  `json:"name"`
	Exchange                 string  `json:"exchange"`
	PeForecast               float64 `json:"pe_forecast"`
	TotalShares              int64   `json:"total_shares"`
	Status                   int     `json:"status"`
	Code                     string  `json:"code"`
	AvgPrice                 float64 `json:"avg_price"`
	Percent                  float64 `json:"percent"`
	Psr                      float64 `json:"psr"`
	Amplitude                float64 `json:"amplitude"`
	Current                  float64 `json:"current"`
	CurrentYearPercent       float64 `json:"current_year_percent"`
	IssueDate                int64   `json:"issue_date"`
	SubType                  string  `json:"sub_type"`
	Low                      float64 `json:"low"`
	MarketCapital            float64 `json:"market_capital"`
	ShareholderFunds         float64 `json:"shareholder_funds"`
	Dividend                 float64 `json:"dividend"`
	DividendYield            float64 `json:"dividend_yield"`
	Currency                 string  `json:"currency"`
	ChgExt                   float64 `json:"chg_ext"`
	Navps                    float64 `json:"navps"`
	Profit                   float64 `json:"profit"`
	Timestamp                int64   `json:"timestamp"`
	PeLyr                    float64 `json:"pe_lyr"`
	Amount                   float64 `json:"amount"`
	Pb                       float64 `json:"pb"`
	DualCounterMappingSymbol string  `json:"dual_counter_mapping_symbol"`
	PeTtm                    float64 `json:"pe_ttm"`
	VariableTickSize         string  `json:"variable_tick_size"`
	Time                     int64   `json:"time"`
	Open                     float64 `json:"open"`
}

type QuoteMarket struct {
	Region string `json:"region"`
	Status string `json:"status"`
}

type QuoteOthers struct {
	PankouRatio float64 `json:"pankou_ratio"`
	CybSwitch   bool    `json:"cyb_switch"`
}

type QuoteTags struct {
	Description string `json:"description"`
	Value       int    `json:"value"`
}
