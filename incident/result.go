package incident

import "github.com/opsgenie/opsgenie-go-sdk-v2/client"

type RequestStatusResult struct {
	client.ResultMetadata
	Success       bool   `json:"success"`
	Action        string `json:"action"`
	ProcessedAt   string `json:"processedAt"`
	IntegrationId string `json:"integrationId"`
	IsSuccess     bool   `json:"isSuccess"`
	Status        string `json:"status"`
	IncidentId    string `json:"incidentId"`
}

type AsyncResult struct {
	client.ResultMetadata
	Result string `json:"result"`
}

type IncidentResult struct {
	client.ResultMetadata
	Id              string            `json:"id"`
	ServiceId       string            `json:"serviceId"`
	TinyId          string            `json:"tinyId"`
	Message         string            `json:"message"`
	Status          string            `json:"status"`
	Tags            []string          `json:"tags"`
	CreatedAt       string            `json:"createdAt"`
	UpdatedAt       string            `json:"updatedAt"`
	Priority        Priority          `json:"priority"`
	OwnerTeam       string            `json:"ownerTeam"`
	Responders      []Responder       `json:"responders"`
	ExtraProperties map[string]string `json:"extraProperties"`
}

type ListResult struct {
	client.ResultMetadata
	Incidents []IncidentResult `json:"data"`
	Paging    Paging           `json:"paging"`
}

type LogResult struct {
	Log       string `json:"log"`
	Type      string `json:"type"`
	Owner     string `json:"owner"`
	CreatedAt string `json:"createdAt"`
	Offset    string `json:"offset"`
}

type ListLogsResult struct {
	client.ResultMetadata
	Logs   []LogResult `json:"data"`
	Paging Paging      `json:"paging"`
}

type NoteResult struct {
	Note      string `json:"note"`
	Owner     string `json:"owner"`
	CreatedAt string `json:"createdAt"`
	Offset    string `json:"offset"`
}

type ListNotesResult struct {
	client.ResultMetadata
	Notes  []NoteResult `json:"data"`
	Paging Paging       `json:"paging"`
}

type Paging struct {
	Next  string `json:"next"`
	First string `json:"first"`
	Last  string `json:"last"`
}
