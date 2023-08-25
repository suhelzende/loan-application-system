package model

type AccountingProvider struct {
	ID   string
	Name string
}

type BalenceSheet struct {
	Sheet []BalanceSheetForMonth
}

type BalanceSheetForMonth struct {
	Year         int
	Month        int
	ProfitOrLoss int
	AssetsValue  int
}
