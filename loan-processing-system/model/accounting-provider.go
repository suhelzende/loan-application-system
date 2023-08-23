package model

type AccountingProvider struct {
	ID   string
	Name string
}

type BalenceSheet struct {
	Sheet []BalenceSheetForMonth
}

type BalenceSheetForMonth struct {
	Year         int
	Month        int
	ProfitOrLoss int
	AssetsValue  int
}
