package models

type Response struct {
    Element_count      int              `json:"element_count"Near_earth_objects"`
    Near_earth_objects map[string][]NEO `json:"near_earth_objects"`
}
