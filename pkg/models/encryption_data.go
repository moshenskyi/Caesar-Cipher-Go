package models

type EncryptionData struct {
	Value string `json:"message"`
	Shift uint8  `json:"shift"`
}
