package request

type Scores struct {
	Seiri    int `json:"SEIRI" validate:"required"`
	Seiton   int `json:"SEITON" validate:"required"`
	Seiso    int `json:"SEISO" validate:"required"`
	Seiketsu int `json:"SEIKETSU" validate:"required"`
	Shitsuke int `json:"SHITSUKE" validate:"required"`
}
