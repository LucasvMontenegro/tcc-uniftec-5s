package datastructure

type Prize struct {
	ID          *int64 `gorm:"primarykey"`
	EditionID   *int64
	Name        *string
	Description *string
}
