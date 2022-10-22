package entity

type FiveS struct {
	ID          *int64
	Name        string
	Description *string
}

type FiveSInterface interface {
	Self() *Score
}

type FiveSRepository interface {
}

type FiveSFactoryInterface interface {
}
