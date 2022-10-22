package datastructure

type Score struct {
	ID      *int64 `gorm:"primarykey"`
	FiveSID *int64
	FiveS   *FiveS
	TeamID  *int64
	Team    *Team
	Score   *float64
}

type FiveS struct {
	ID   *int64
	Name *string
}
