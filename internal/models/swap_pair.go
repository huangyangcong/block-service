package models

import "strings"

type Currency struct {
	Contract string
	Symbol   string
	Decimal  int
	Id       string
}

// A->B(A兑换为B)
type CurrencyPair struct {
	CurrencyA      Currency
	CurrencyB      Currency
	PairId         string // 交易对ID
	AmountTickSize int    // 下单量精度
	PriceTickSize  int    //交易对价格精度
}

var (
	UNKNOWN = Currency{"UNKNOWN", "", 0, ""}
	EOS     = Currency{"eosio.token", "EOS", 4, ""}
	BOX     = Currency{"token.defi", "BOX", 4, ""}
	USDT    = Currency{"tethertether", "USDT", 4, ""}

	//currency pair
	BOX_EOS      = CurrencyPair{CurrencyA: BOX, CurrencyB: EOS, PairId: "194", AmountTickSize: 4, PriceTickSize: 4}
	USDT_EOS     = CurrencyPair{CurrencyA: USDT, CurrencyB: EOS, PairId: "12", AmountTickSize: 4, PriceTickSize: 4}
	UNKNOWN_PAIR = CurrencyPair{CurrencyA: UNKNOWN, CurrencyB: UNKNOWN}
)

func (pair CurrencyPair) ToSymbol(joinChar string) string {
	return strings.Join([]string{pair.CurrencyA.Symbol, pair.CurrencyB.Symbol}, joinChar)
}
