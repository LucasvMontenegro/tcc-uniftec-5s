package entity

type Team struct {
}

type TeamInterface interface {
}

type TeamRepository interface {
}

type TeamFactoryInterface interface {
	NewTeam() TeamInterface
}
