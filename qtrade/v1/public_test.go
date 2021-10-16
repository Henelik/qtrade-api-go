package qtrade

import (
	"context"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	commonTestData = `{"data": {"currencies": [{"can_withdraw": false,"code": "MMO","config": {"address_version": 50,"default_signer": 23,"price": 0.002,"required_confirmations": 6,"required_generate_confirmations": 120,"satoshi_per_byte": 100,"wif_version": 178,"withdraw_fee": "0.001"},"long_name": "MMOCoin","metadata": {"delisting_date": "12/13/2018"},"precision": 8,"status": "delisted","type": "bitcoin_like"},{"can_withdraw": true,"code": "BTC","config": {"address_version": 0,"default_signer": 6,"explorerAddressURL": "https://live.blockcypher.com/btc/address/","explorerTransactionURL": "https://live.blockcypher.com/btc/tx/","p2sh_address_version": 5,"price": 8595.59,"required_confirmations": 2,"required_generate_confirmations": 100,"satoshi_per_byte": 15,"withdraw_fee": "0.0005"},"long_name": "Bitcoin","metadata": {"withdraw_notices": []},"precision": 8,"status": "ok","type": "bitcoin_like"},{"can_withdraw": true,"code": "BIS","config": {"data_max": 1000,"default_signer": 54,"enable_address_data": true,"explorerAddressURL": "https://bismuth.online/search?quicksearch=","explorerTransactionURL": "https://bismuth.online/search?quicksearch=","price": 0.12891653720125376,"required_confirmations": 35,"withdraw_fee": "0.25"},"long_name": "Bismuth","metadata": {"deposit_notices": [],"hidden": false},"precision": 8,"status": "ok","type": "bismuth"}],"markets": [{"base_currency": "BTC","can_cancel": false,"can_trade": false,"can_view": false,"id": 8,"maker_fee": "0.0025","market_currency": "MMO","metadata": {"delisting_date": "12/13/2018","market_notices": [{"message": "Delisting Notice: This market has been delisted due to low volume. Please cancel your orders and withdraw your funds by 12/13/2018.","type": "warning"}]},"taker_fee": "0.0025"},{"base_currency": "BTC","can_cancel": true,"can_trade": true,"can_view": true,"id": 20,"maker_fee": "0","market_currency": "BIS","metadata": {"labels": []},"taker_fee": "0.005"}],"tickers": [{"ask": null,"bid": null,"day_avg_price": null,"day_change": null,"day_high": null,"day_low": null,"day_open": null,"day_volume_base": "0","day_volume_market": "0","id": 8,"id_hr": "MMO_BTC","last": "0.00000076"},{"ask": "0.000014","bid": "0.00001324","day_avg_price": "0.0000147000191353","day_change": "-0.023086269744836","day_high": "0.00001641","day_low": "0.00001292","day_open": "0.00001646","day_volume_base": "0.36885974","day_volume_market": "25092.46665642","id": 20,"id_hr": "BIS_BTC","last": "0.00001608"}]}}`
)

func TestClient_GetCommon(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "http://localhost/v1/common",
		httpmock.NewStringResponder(200, commonTestData))

	want := &CommonData{
		Currencies: []CurrencyData{
			{
				CanWithdraw: false,
				Code:        MMO,
				Config: CurrencyConfig{
					AddressVersion:                50,
					DefaultSigner:                 23,
					Price:                         0.002,
					RequiredConfirmations:         6,
					RequiredGenerateConfirmations: 120,
					SatoshiPerByte:                100,
					WifVersion:                    178,
					WithdrawFee:                   "0.001",
					ExplorerAddressURL:            "",
					ExplorerTransactionURL:        "",
					P2ShAddressVersion:            0,
					DataMax:                       0,
					EnableAddressData:             false,
				},
				LongName: "MMOCoin",
				Metadata: CurrencyMetadata{
					DelistingDate:   "12/13/2018",
					WithdrawNotices: nil,
					DepositNotices:  nil,
					Hidden:          false},
				Precision: 8,
				Status:    "delisted",
				Type:      "bitcoin_like",
			},
			{
				CanWithdraw: true,
				Code:        BTC,
				Config: CurrencyConfig{
					AddressVersion:                0,
					DefaultSigner:                 6,
					Price:                         8595.59,
					RequiredConfirmations:         2,
					RequiredGenerateConfirmations: 100,
					SatoshiPerByte:                15,
					WifVersion:                    0,
					WithdrawFee:                   "0.0005",
					ExplorerAddressURL:            "https://live.blockcypher.com/btc/address/",
					ExplorerTransactionURL:        "https://live.blockcypher.com/btc/tx/",
					P2ShAddressVersion:            5,
					DataMax:                       0,
					EnableAddressData:             false,
				},
				LongName: "Bitcoin",
				Metadata: CurrencyMetadata{
					DelistingDate:   "",
					WithdrawNotices: []interface{}{},
					DepositNotices:  nil,
					Hidden:          false,
				},
				Precision: 8,
				Status:    "ok",
				Type:      "bitcoin_like",
			},
			{
				CanWithdraw: true,
				Code:        BIS,
				Config: CurrencyConfig{
					AddressVersion:                0,
					DefaultSigner:                 54,
					Price:                         0.12891653720125376,
					RequiredConfirmations:         35,
					RequiredGenerateConfirmations: 0,
					SatoshiPerByte:                0,
					WifVersion:                    0,
					WithdrawFee:                   "0.25",
					ExplorerAddressURL:            "https://bismuth.online/search?quicksearch=",
					ExplorerTransactionURL:        "https://bismuth.online/search?quicksearch=",
					P2ShAddressVersion:            0,
					DataMax:                       1000,
					EnableAddressData:             true,
				},
				LongName: "Bismuth",
				Metadata: CurrencyMetadata{
					DelistingDate:   "",
					WithdrawNotices: nil,
					DepositNotices:  []interface{}{},
					Hidden:          false,
				},
				Precision: 8,
				Status:    "ok",
				Type:      "bismuth",
			},
		},
		Markets: []MarketData{
			{
				BaseCurrency:   BTC,
				CanCancel:      false,
				CanTrade:       false,
				CanView:        false,
				Market:         MMO_BTC,
				MakerFee:       0.0025,
				MarketCurrency: MMO,
				Metadata: MarketMetadata{
					DelistingDate: "12/13/2018",
					MarketNotices: []MarketNotice{
						{
							Message: "Delisting Notice: This market has been delisted due to low volume. Please cancel your orders and withdraw your funds by 12/13/2018.",
							Type:    "warning"},
					},
					Labels: nil,
				},
				TakerFee: 0.0025,
			},
			{
				BaseCurrency:   "BTC",
				CanCancel:      true,
				CanTrade:       true,
				CanView:        true,
				Market:         BIS_BTC,
				MakerFee:       0,
				MarketCurrency: "BIS",
				Metadata: MarketMetadata{
					DelistingDate: "",
					MarketNotices: nil,
					Labels:        []interface{}{},
				},
				TakerFee: 0.005,
			},
		},
		Tickers: []Ticker{
			{
				Ask:             0,
				Bid:             0,
				DayAvgPrice:     0,
				DayChange:       0,
				DayHigh:         0,
				DayLow:          0,
				DayOpen:         0,
				DayVolumeBase:   0,
				DayVolumeMarket: 0,
				Market:          MMO_BTC,
				IdHr:            "MMO_BTC",
				Last:            0.00000076,
			},
			{
				Ask:             0.000014,
				Bid:             0.00001324,
				DayAvgPrice:     0.0000147000191353,
				DayChange:       -0.023086269744836,
				DayHigh:         0.00001641,
				DayLow:          0.00001292,
				DayOpen:         0.00001646,
				DayVolumeBase:   0.36885974,
				DayVolumeMarket: 25092.46665642,
				Market:          BIS_BTC,
				IdHr:            "BIS_BTC",
				Last:            0.00001608,
			},
		},
	}

	got, err := testClient.GetCommon(context.Background())
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}

	assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET http://localhost/v1/common"])
}
