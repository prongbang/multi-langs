package model

type Message struct {
	DeviceID int    `json:"deviceId"`
	Status   string `json:"status"`
	Topic    string `json:"topic"`
	Value    string `json:"value"`
}
