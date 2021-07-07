package schedule

import (
	"block-service/internal/models"
	"block-service/internal/util"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nntaoli-project/goex"
	"github.com/nntaoli-project/goex/okex"
	"net/http"
	"net/url"
	"strings"
)

var BOX_USDT = goex.CurrencyPair{CurrencyA: goex.Currency{"BOX", ""}, CurrencyB: goex.USDT, AmountTickSize: 4, PriceTickSize: 4}

type BoxPrice struct {
}

func NewBoxPrice(s *Server, m *util.EmailNotify) BoxPrice {
	log := log.NewHelper(s.log)
	s.schedule.AddFunc("@every 3s", func() {
		err, boxEos := util.GetSwapDefiPair(models.BOX_EOS.PairId)
		err, usdtEos := util.GetSwapDefiPair(models.USDT_EOS.PairId)
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
		var okex = okex.NewOKEx(&goex.APIConfig{
			Endpoint: "https://www.okex.com",
			HttpClient: &http.Client{
				Transport: &http.Transport{
					Proxy: func(req *http.Request) (*url.URL, error) {
						return &url.URL{
							Scheme: "socks5",
							Host:   "127.0.0.1:1080"}, nil
					},
				},
			},
			ApiKey:        "",
			ApiSecretKey:  "",
			ApiPassphrase: "",
		})
		var okexSpot = okex.OKExSpot
		ticker, err := okexSpot.GetTicker(BOX_USDT)
		if boxUsdtPrice-ticker.Last > 5 {
			sr := fmt.Sprintf("box defibox价格为：%f okex价格为：%f", boxUsdtPrice, ticker.Last)
			m.SendNotifyWithFile("xxxx", "box价格监控", sr)
		}
	})
	return BoxPrice{}
}
