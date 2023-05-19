package models

type Environment struct {
	Name          string `json:"name"`
	UUID          string `json:"uuid"`
	OwnerUUID     string `json:"owner_uuid"`
	PrototypeUUID string `json:"prototype_uuid"`
}
