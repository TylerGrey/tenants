package model

import "strconv"

// KakaoResponse ...
type KakaoResponse struct {
	Documents []KakaoDocument `json:"documents"`
}

// KakaoDocument 카카오 주소 검색 Response
type KakaoDocument struct {
	Address     *KakaoAddress     `json:"address"`
	RoadAddress *KakaoRoadAddress `json:"road_address"`
	X           string            `json:"x"`
	Y           string            `json:"y"`
}

// KakaoAddress 카카오 주소 검색 Response
type KakaoAddress struct {
	AddressName string `json:"address_name"`
}

// KakaoRoadAddress 카카오 주소 검색 Response
type KakaoRoadAddress struct {
	AddressName  string `json:"address_name"`
	BuildingName string `json:"building_name"`
	RoadName     string `json:"road_name"`
	ZoneNo       string `json:"zone_no"`
}

// SearchResult ...
type SearchResult struct {
	Payload KakaoDocument
}

// Name ...
func (sr SearchResult) Name() string {
	var name string
	if sr.Payload.RoadAddress.BuildingName != "" {
		name = sr.Payload.RoadAddress.BuildingName
	} else {
		name = sr.Payload.RoadAddress.RoadName
	}
	return name
}

// Address ...
func (sr SearchResult) Address() string {
	return sr.Payload.Address.AddressName
}

// RoadAddress ...
func (sr SearchResult) RoadAddress() string {
	return sr.Payload.RoadAddress.AddressName
}

// ZoneNo ...
func (sr SearchResult) ZoneNo() string {
	return sr.Payload.RoadAddress.ZoneNo
}

// Lat ...
func (sr SearchResult) Lat() float64 {
	lat, err := strconv.ParseFloat(sr.Payload.Y, 64)
	if err != nil {
		lat = 0
	}

	return lat
}

// Lng ...
func (sr SearchResult) Lng() float64 {
	lng, err := strconv.ParseFloat(sr.Payload.X, 64)
	if err != nil {
		lng = 0
	}

	return lng
}
