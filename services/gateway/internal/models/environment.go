package models

type Environment struct {
	Type          string `json:"type"`
	EnvIdentifier string `json:"env_identifier"`
	Name          string `json:"name"`
}
