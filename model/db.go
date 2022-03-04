package model

type DbUpdateMember struct {
	UserId  int            `json:"user_id"`
	Updates map[string]any `json:"updates"`
}
