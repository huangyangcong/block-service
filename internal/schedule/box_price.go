package schedule

import (
	"encoding/json"
	"github.com/eoscanada/eos-go"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nntaoli-project/goex"
	okex2 "github.com/nntaoli-project/goex/okex"
	"github.com/robfig/cron/v3"
	"net/http"
	"strings"
)

var BOX_USDT = goex.CurrencyPair{CurrencyA: goex.Currency{"BOX", ""}, CurrencyB: goex.USDT, AmountTickSize: 4, PriceTickSize: 4}

type pair struct {
	Price1CumulativeLast string  `json:"price1_cumulative_last" gorm:"column:price1_cumulative_last"`
	BlockTimeLast        string  `json:"block_time_last" gorm:"column:block_time_last"`
	Price0Last           float64 `json:"price0_last,string" gorm:"column:price0_last"`
	Price1Last           float64 `json:"price1_last,string" gorm:"column:price1_last"`
	Token0               struct {
		Symbol   string `json:"symbol" gorm:"column:symbol"`
		Contract string `json:"contract" gorm:"column:contract"`
	} `json:"token0" gorm:"column:token0"`
	LiquidityToken int `json:"liquidity_token" gorm:"column:liquidity_token"`
	Token1         struct {
		Symbol   string `json:"symbol" gorm:"column:symbol"`
		Contract string `json:"contract" gorm:"column:contract"`
	} `json:"token1" gorm:"column:token1"`
	Reserve1             string `json:"reserve1" gorm:"column:reserve1"`
	Reserve0             string `json:"reserve0" gorm:"column:reserve0"`
	ID                   int    `json:"id" gorm:"column:id"`
	Price0CumulativeLast int    `json:"price0_cumulative_last" gorm:"column:price0_cumulative_last"`
}

type marketTicket struct {
	Msg  string `json:"msg" gorm:"column:msg"`
	Code string `json:"code" gorm:"column:code"`
	Data []struct {
		Last      float64 `json:"last,string" gorm:"column:last"`
		LastSz    float64 `json:"lastSz,string" gorm:"column:lastSz"`
		Open24h   float64 `json:"open24h,string" gorm:"column:open24h"`
		AskSz     float64 `json:"askSz,string" gorm:"column:askSz"`
		Low24h    float64 `json:"low24h,string" gorm:"column:low24h"`
		AskPx     float64 `json:"askPx,string" gorm:"column:askPx"`
		VolCcy24h float64 `json:"volCcy24h,string" gorm:"column:volCcy24h"`
		InstType  string  `json:"instType" gorm:"column:instType"`
		InstID    string  `json:"instId" gorm:"column:instId"`
		BidSz     float64 `json:"bidSz,string" gorm:"column:bidSz"`
		BidPx     float64 `json:"bidPx,string" gorm:"column:bidPx"`
		High24h   float64 `json:"high24h,string" gorm:"column:high24h"`
		SodUtc0   float64 `json:"sodUtc0,string" gorm:"column:sodUtc0"`
		Vol24h    float64 `json:"vol24h,string" gorm:"column:vol24h"`
		SodUtc8   float64 `json:"sodUtc8,string" gorm:"column:sodUtc8"`
		Ts        uint    `json:"ts,string" gorm:"column:ts"`
	} `json:"data" gorm:"column:data"`
}

type BoxPrice struct {
}

func NewBoxPrice(c *cron.Cron, logger log.Logger) BoxPrice {
	log := log.NewHelper(logger)
	c.AddFunc("@every 3s", func() {
		err, boxEos := getPair("194")
		err, usdtEos := getPair("12")
		var (
			boxEosPrice, usdtEosPrice float64
		)
		if strings.HasSuffix(boxEos.Token0.Symbol, "EOS") {
			boxEosPrice = boxEos.Price1Last
		} else {
			boxEosPrice = boxEos.Price0Last
		}
		if strings.HasSuffix(usdtEos.Token0.Symbol, "EOS") {
			usdtEosPrice = usdtEos.Price0Last
		} else {
			usdtEosPrice = usdtEos.Price1Last
		}
		boxUsdtPrice := boxEosPrice * usdtEosPrice
		if err == nil {
			log.Infof("defibox boxUsdtPrice=%f", boxUsdtPrice)
		} else {
			log.Error(err)
		}
		// Create a Resty Client
		var okex = okex2.NewOKEx(&goex.APIConfig{
			Endpoint:   "https://www.okex.com",
			HttpClient: &http.Client{
				//Transport: &http.Transport{
				//
				//},
			},
			ApiKey:        "",
			ApiSecretKey:  "",
			ApiPassphrase: "",
		})
		var (
			okexSpot = okex.OKExSpot
			//okexSwap   = okex.OKExSwap   //永续合约实现
			//okexFuture = okex.OKExFuture //交割合约实现
			//okexWallet = okex.OKExWallet //资金账户（钱包）操作
		)
		ticker, err := okexSpot.GetTicker(BOX_USDT)
		log.Info(ticker)
	})
	return BoxPrice{}
}

func getPair(pairId string) (error, pair) {
	api := eos.New("https://eospush.tokenpocket.pro")
	boxEos, err := api.GetTableRows(eos.GetTableRowsRequest{
		Code:       "swap.defi",
		Scope:      "swap.defi",
		Table:      "pairs",
		JSON:       true,
		Limit:      1000,
		LowerBound: pairId,
		UpperBound: pairId,
	})
	var pairs []pair
	json.Unmarshal(boxEos.Rows, &pairs)
	return err, pairs[0]
}
