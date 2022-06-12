package models

type NEO struct {
    Id        string `json:"id"`
    Name      string `json:"name"`
    Hazardous bool   `json:"is_potentially_hazardous_asteroid"`
}
