package models

type Environment struct {
	Name          string `json:"name"`
	UUID          string `json:"uuid"`
	OwnerUUID     string `json:"owner_uuid"`
	PrototypeUUID string `json:"prototype_uuid"`
}

type CreateEnvironmentRequest struct {
	PrototypeUUID string `json:"prototype_uuid"`
	Name          string `json:"name"`
}

type CreateEnvironmentResponse struct {
	EnvironmentUUID string `json:"environment_uuid"`
}
