package models

import (
	"encoding/json"
)

type Source struct {
	Provider    string          `json:"provider"`
	CalendarId  string          `json:"calendar_id"`
	EventName   string          `json:"event_name"`
	Credentials json.RawMessage `json:"credentials"`
}

type Version struct {
	Id string `json:"id"`
}

type CheckRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

type InRequest struct {
	Source  Source          `json:"source"`
	Version Version         `json:"version"`
	Params  json.RawMessage `json:"params"`
}

type InResponse struct {
	Version  Version        `json:"version"`
	MetaData []KeyValuePair `json:"metadata"`
}

type KeyValuePair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type OutRequest struct {
	Source Source          `json:"source"`
	Params json.RawMessage `json:"params"`
}

type OutResponse struct {
	Version Version `json:"version"`
}
