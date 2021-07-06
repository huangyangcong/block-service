package util

import (
	"encoding/json"
	"github.com/eoscanada/eos-go"
)

type SwapDefiPair struct {
	ID int `json:"id" gorm:"column:id"`

	Token0 struct {
		Symbol   string `json:"symbol" gorm:"column:symbol"`
		Contract string `json:"contract" gorm:"column:contract"`
	} `json:"token0" gorm:"column:token0"`
	Token1 struct {
		Symbol   string `json:"symbol" gorm:"column:symbol"`
		Contract string `json:"contract" gorm:"column:contract"`
	} `json:"token1" gorm:"column:token1"`
	BlockTimeLast        string  `json:"block_time_last" gorm:"column:block_time_last"`
	LiquidityToken       int     `json:"liquidity_token,string" gorm:"column:liquidity_token"`
	Reserve1             string  `json:"reserve1" gorm:"column:reserve1"`
	Reserve0             string  `json:"reserve0" gorm:"column:reserve0"`
	Price0Last           float64 `json:"price0_last,string" gorm:"column:price0_last"`
	Price1Last           float64 `json:"price1_last,string" gorm:"column:price1_last"`
	Price0CumulativeLast int     `json:"price0_cumulative_last" gorm:"column:price0_cumulative_last"`
	Price1CumulativeLast int     `json:"price1_cumulative_last" gorm:"column:price1_cumulative_last"`
}

func GetSwapDefiPair(pairId string) (error, *SwapDefiPair) {
	var pairs []SwapDefiPair
	err := GetTableRows("swap.defi", "swap.defi", "pairs", pairId, pairId, 1000, &pairs)
	return err, &pairs[0]
}

// GetTableRows
func GetTableRows(code, scope, table, lowerBound, upperBound string, limit uint32, v interface{}) error {
	api := eos.New("https://eospush.tokenpocket.pro")
	resp, err := api.GetTableRows(eos.GetTableRowsRequest{
		Code:       code,
		Scope:      scope,
		Table:      table,
		JSON:       true,
		Limit:      limit,
		LowerBound: lowerBound,
		UpperBound: upperBound,
	})
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp.Rows, &v)
	if err != nil {
		return err
	}
	return nil

}
