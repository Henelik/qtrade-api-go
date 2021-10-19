package qtrade

import (
	"context"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	commonTestData     = `{"data": {"currencies": [{"can_withdraw": false,"code": "MMO","config": {"address_version": 50,"default_signer": 23,"price": 0.002,"required_confirmations": 6,"required_generate_confirmations": 120,"satoshi_per_byte": 100,"wif_version": 178,"withdraw_fee": "0.001"},"long_name": "MMOCoin","metadata": {"delisting_date": "12/13/2018"},"precision": 8,"status": "delisted","type": "bitcoin_like"},{"can_withdraw": true,"code": "BTC","config": {"address_version": 0,"default_signer": 6,"explorerAddressURL": "https://live.blockcypher.com/btc/address/","explorerTransactionURL": "https://live.blockcypher.com/btc/tx/","p2sh_address_version": 5,"price": 8595.59,"required_confirmations": 2,"required_generate_confirmations": 100,"satoshi_per_byte": 15,"withdraw_fee": "0.0005"},"long_name": "Bitcoin","metadata": {"withdraw_notices": []},"precision": 8,"status": "ok","type": "bitcoin_like"},{"can_withdraw": true,"code": "BIS","config": {"data_max": 1000,"default_signer": 54,"enable_address_data": true,"explorerAddressURL": "https://bismuth.online/search?quicksearch=","explorerTransactionURL": "https://bismuth.online/search?quicksearch=","price": 0.12891653720125376,"required_confirmations": 35,"withdraw_fee": "0.25"},"long_name": "Bismuth","metadata": {"deposit_notices": [],"hidden": false},"precision": 8,"status": "ok","type": "bismuth"}],"markets": [{"base_currency": "BTC","can_cancel": false,"can_trade": false,"can_view": false,"id": 8,"maker_fee": "0.0025","market_currency": "MMO","metadata": {"delisting_date": "12/13/2018","market_notices": [{"message": "Delisting Notice: This market has been delisted due to low volume. Please cancel your orders and withdraw your funds by 12/13/2018.","type": "warning"}]},"taker_fee": "0.0025"},{"base_currency": "BTC","can_cancel": true,"can_trade": true,"can_view": true,"id": 20,"maker_fee": "0","market_currency": "BIS","metadata": {"labels": []},"taker_fee": "0.005"}],"tickers": [{"ask": null,"bid": null,"day_avg_price": null,"day_change": null,"day_high": null,"day_low": null,"day_open": null,"day_volume_base": "0","day_volume_market": "0","id": 8,"id_hr": "MMO_BTC","last": "0.00000076"},{"ask": "0.000014","bid": "0.00001324","day_avg_price": "0.0000147000191353","day_change": "-0.023086269744836","day_high": "0.00001641","day_low": "0.00001292","day_open": "0.00001646","day_volume_base": "0.36885974","day_volume_market": "25092.46665642","id": 20,"id_hr": "BIS_BTC","last": "0.00001608"}]}}`
	tickerTestData     = `{"data": {"ask": "0.02249","bid": "0.0191","day_avg_price": "0.0197095311101552","day_change": "0.0380429141071119","day_high": "0.02249","day_low": "0.0184","day_open": "0.01840001","day_volume_base": "0.42644484","day_volume_market": "21.63647819","id": 15,"id_hr": "VEO_BTC","last": "0.0191"}}`
	tickersTestData    = `{"data": {"markets": [{"ask": "0.0034","bid": "0.0011","day_avg_price": null,"day_change": null,"day_high": null,"day_low": null,"day_open": null,"day_volume_base": "0","day_volume_market": "0","id": 23,"id_hr": "GRIN_BTC","last": "0.0033"},{"ask": "0.000099","bid": "0.0000795","day_avg_price": "0.0000795337894515","day_change": "-0.2205882352941176","day_high": "0.00008","day_low": "0.0000795","day_open": "0.000102","day_volume_base": "0.07353291","day_volume_market": "924.549308","id": 19,"id_hr": "SNOW_BTC","last": "0.0000795"}]}}`
	currencyTestData   = `{"data": {"currency": {"can_withdraw": true,"code": "BTC","config": {"address_version": 0,"default_signer": 6,"explorerAddressURL": "https://live.blockcypher.com/btc/address/","explorerTransactionURL": "https://live.blockcypher.com/btc/tx/","p2sh_address_version": 5,"price": 9159.72,"required_confirmations": 2,"required_generate_confirmations": 100,"satoshi_per_byte": 15,"withdraw_fee": "0.0005"},"long_name": "Bitcoin","metadata": {"withdraw_notices": []},"precision": 8,"status": "ok","type": "bitcoin_like"}}}`
	currenciesTestData = `{"data": {"currencies": [{"can_withdraw": true,"code": "BTC","config": {"address_version": 0,"default_signer": 6,"explorerAddressURL": "https://live.blockcypher.com/btc/address/","explorerTransactionURL": "https://live.blockcypher.com/btc/tx/","p2sh_address_version": 5,"price": 9159.72,"required_confirmations": 2,"required_generate_confirmations": 100,"satoshi_per_byte": 15,"withdraw_fee": "0.0005"},"long_name": "Bitcoin","metadata": {"withdraw_notices": []},"precision": 8,"status": "ok","type": "bitcoin_like"},{"can_withdraw": true,"code": "BIS","config": {"data_max": 1000,"default_signer": 54,"enable_address_data": true,"explorerAddressURL": "https://bismuth.online/search?quicksearch=","explorerTransactionURL": "https://bismuth.online/search?quicksearch=","price": 0.11314929085578249,"required_confirmations": 35,"withdraw_fee": "0.25"},"long_name": "Bismuth","metadata": {"deposit_notices": [],"hidden": false},"precision": 8,"status": "ok","type": "bismuth"}]}}`
	marketTestData     = `{"data": {"market": {"base_currency": "BTC","can_cancel": true,"can_trade": true,"can_view": false,"id": 15,"maker_fee": "0.005","market_currency": "VEO","metadata": {},"taker_fee": "0.005"},"recent_trades": [{"amount": "1.64360163","created_at": "2019-01-31T23:09:31.419131Z","id": 51362,"price": "0.0191"},{"amount": "1.60828469","created_at": "2019-01-31T22:05:16.531659Z","id": 51362,"price": "0.02248"}]}}`
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
					WithdrawFee:                   0.001,
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
					WithdrawFee:                   0.0005,
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
					WithdrawFee:                   0.25,
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
				ID:             MMO_BTC,
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
				ID:             BIS_BTC,
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

func TestClient_GetTicker(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "http://localhost/v1/ticker/VEO_BTC",
		httpmock.NewStringResponder(200, tickerTestData))

	want := &Ticker{
		Ask:             0.02249,
		Bid:             0.0191,
		DayAvgPrice:     0.0197095311101552,
		DayChange:       0.0380429141071119,
		DayHigh:         0.02249,
		DayLow:          0.0184,
		DayOpen:         0.01840001,
		DayVolumeBase:   0.42644484,
		DayVolumeMarket: 21.63647819,
		Market:          VEO_BTC,
		IdHr:            VEO_BTC.String(),
		Last:            0.0191,
	}

	got, err := testClient.GetTicker(context.Background(), VEO_BTC)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}

	assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET http://localhost/v1/ticker/VEO_BTC"])
}

func TestClient_GetTickers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "http://localhost/v1/tickers",
		httpmock.NewStringResponder(200, tickersTestData))

	want := []Ticker{
		{
			Ask:             0.0034,
			Bid:             0.0011,
			DayAvgPrice:     0,
			DayChange:       0,
			DayHigh:         0,
			DayLow:          0,
			DayOpen:         0,
			DayVolumeBase:   0,
			DayVolumeMarket: 0,
			Market:          GRIN_BTC,
			IdHr:            "GRIN_BTC",
			Last:            0.0033,
		},
		{
			Ask:             0.000099,
			Bid:             0.0000795,
			DayAvgPrice:     0.0000795337894515,
			DayChange:       -0.2205882352941176,
			DayHigh:         0.00008,
			DayLow:          0.0000795,
			DayOpen:         0.000102,
			DayVolumeBase:   0.07353291,
			DayVolumeMarket: 924.549308,
			Market:          SNOW_BTC,
			IdHr:            "SNOW_BTC",
			Last:            0.0000795,
		},
	}

	got, err := testClient.GetTickers(context.Background())
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}

	assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET http://localhost/v1/tickers"])
}

func TestClient_GetCurrency(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "http://localhost/v1/currency/BTC",
		httpmock.NewStringResponder(200, currencyTestData))

	want := &CurrencyData{
		CanWithdraw: true,
		Code:        BTC,
		Config: CurrencyConfig{
			AddressVersion:                0,
			DefaultSigner:                 6,
			Price:                         9159.72,
			RequiredConfirmations:         2,
			RequiredGenerateConfirmations: 100,
			SatoshiPerByte:                15,
			WithdrawFee:                   0.0005,
			ExplorerAddressURL:            "https://live.blockcypher.com/btc/address/",
			ExplorerTransactionURL:        "https://live.blockcypher.com/btc/tx/",
			P2ShAddressVersion:            5,
		},
		LongName: "Bitcoin",
		Metadata: CurrencyMetadata{
			WithdrawNotices: []interface{}{},
		},
		Precision: 8,
		Status:    "ok",
		Type:      "bitcoin_like",
	}

	got, err := testClient.GetCurrency(context.Background(), BTC)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}

	assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET http://localhost/v1/currency/BTC"])
}

func TestClient_GetCurrencies(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "http://localhost/v1/currencies",
		httpmock.NewStringResponder(200, currenciesTestData))

	want := []CurrencyData{
		{
			CanWithdraw: true,
			Code:        BTC,
			Config: CurrencyConfig{
				AddressVersion:                0,
				DefaultSigner:                 6,
				Price:                         9159.72,
				RequiredConfirmations:         2,
				RequiredGenerateConfirmations: 100,
				SatoshiPerByte:                15,
				WithdrawFee:                   0.0005,
				ExplorerAddressURL:            "https://live.blockcypher.com/btc/address/",
				ExplorerTransactionURL:        "https://live.blockcypher.com/btc/tx/",
				P2ShAddressVersion:            5,
			},
			LongName: "Bitcoin",
			Metadata: CurrencyMetadata{
				WithdrawNotices: []interface{}{},
			},
			Precision: 8,
			Status:    "ok",
			Type:      "bitcoin_like",
		},
		{
			CanWithdraw: true,
			Code:        BIS,
			Config: CurrencyConfig{
				DefaultSigner:          54,
				Price:                  0.11314929085578249,
				RequiredConfirmations:  35,
				WithdrawFee:            0.25,
				ExplorerAddressURL:     "https://bismuth.online/search?quicksearch=",
				ExplorerTransactionURL: "https://bismuth.online/search?quicksearch=",
				EnableAddressData:      true,
				DataMax:                1000,
			},
			LongName: "Bismuth",
			Metadata: CurrencyMetadata{
				DepositNotices: []interface{}{},
				Hidden:         false,
			},
			Precision: 8,
			Status:    "ok",
			Type:      "bismuth",
		},
	}

	got, err := testClient.GetCurrencies(context.Background())
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}

	assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET http://localhost/v1/currencies"])
}

func TestClient_GetMarket(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "http://localhost/v1/market/VEO_BTC",
		httpmock.NewStringResponder(200, marketTestData))

	wantTime1, _ := time.Parse(time.RFC3339Nano, "2019-01-31T23:09:31.419131Z")
	wantTime2, _ := time.Parse(time.RFC3339Nano, "2019-01-31T22:05:16.531659Z")

	want := &GetMarketData{
		Market: MarketData{
			BaseCurrency:   BTC,
			CanCancel:      true,
			CanTrade:       true,
			CanView:        false,
			ID:             VEO_BTC,
			MakerFee:       0.005,
			MarketCurrency: VEO,
			Metadata:       MarketMetadata{},
			TakerFee:       0.005,
		},
		RecentTrades: []PublicTrade{
			{
				Amount:    1.64360163,
				CreatedAt: wantTime1,
				ID:        51362,
				Price:     0.0191,
			},
			{
				Amount:    1.60828469,
				CreatedAt: wantTime2,
				ID:        51362,
				Price:     0.02248,
			},
		},
	}

	got, err := testClient.GetMarket(context.Background(), VEO_BTC)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}

	assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET http://localhost/v1/market/VEO_BTC"])
}
