package datastructure

type Score struct {
	ID      *int64 `gorm:"primarykey"`
	FiveSID *int64
	FiveS   *FiveS
	TeamID  *int64
	Team    *Team
	Score   *int
}

type FiveS struct {
	ID   *int64
	Name *string
}
