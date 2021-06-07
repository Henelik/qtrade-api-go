package v1

// base types

type Balance struct {
	Currency string
	Balance  string
}

// API results

type GetBalancesResult struct {
	Data Balances
}

type Balances struct {
	Balances []Balance
}
