package schedule

import (
	"encoding/json"
	"github.com/eoscanada/eos-go"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
	"strconv"
	"strings"
)

type Pair struct {
	Price1CumulativeLast string `json:"price1_cumulative_last" gorm:"column:price1_cumulative_last"`
	BlockTimeLast        string `json:"block_time_last" gorm:"column:block_time_last"`
	Price1Last           string `json:"price1_last" gorm:"column:price1_last"`
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
	Price0Last           string `json:"price0_last" gorm:"column:price0_last"`
}

type BoxPrice struct {
}

func NewBoxPrice(c *cron.Cron, logger log.Logger) BoxPrice {
	c.AddFunc("@every 1s", func() {

		err, boxEos := getPair("194")
		err, usdtEos := getPair("12")
		var (
			boxEosPrice, usdtEosPrice float64
		)
		if strings.HasSuffix(boxEos.Token0.Symbol, "EOS") {
			boxEosPrice, _ = strconv.ParseFloat(boxEos.Price1Last, 64)
		} else {
			boxEosPrice, _ = strconv.ParseFloat(boxEos.Price0Last, 64)
		}
		if strings.HasSuffix(usdtEos.Token0.Symbol, "EOS") {
			usdtEosPrice, _ = strconv.ParseFloat(usdtEos.Price0Last, 64)
		} else {
			usdtEosPrice, _ = strconv.ParseFloat(usdtEos.Price1Last, 64)
		}

		boxUsdtPrice := boxEosPrice * usdtEosPrice

		logger.Log(log.LevelInfo, boxUsdtPrice, err)
	})
	return BoxPrice{}
}

func getPair(pairId string) (error, Pair) {
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
	var pairs []Pair
	json.Unmarshal(boxEos.Rows, &pairs)
	return err, pairs[0]
}
