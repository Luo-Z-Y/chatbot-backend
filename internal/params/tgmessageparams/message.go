package tgmessageparams

type MessageParams struct {
	Message string `json:"message" validate:"required"`
}
