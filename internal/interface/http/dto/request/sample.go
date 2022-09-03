package request

type CreateSample struct {
	ReferenceUUID string `json:"reference_uuid" validate:"required,uuid"`
}
