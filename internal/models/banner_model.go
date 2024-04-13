package models

import (
	"encoding/json"
	"time"

	"github.com/9Neechan/AvitoTech-2024/internal/database"
)

type Banner struct {
	Feature     int32           `json:"feature"`
	Tag         int32           `json:"tag"`
	JsonContent json.RawMessage `json:"json_content"`
	IsActive    bool            `json: is_active`
	CreatedAt   time.Time       `json:"updated_at"`
	UpdatedAt   time.Time       `json:"name"`
}

func DatabaseBannerToBanner(banner database.Banner) Banner {
	return Banner{
		Feature:     banner.Feature,
		Tag:         banner.Tag,
		JsonContent: banner.JsonContent,
		IsActive:    banner.IsActive,
		CreatedAt:   banner.CreatedAt,
		UpdatedAt:   banner.UpdatedAt,
	}
}

func UnmarshalJson(jsonData string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
func parseJSONFile(jsonFile []byte) (map[string]interface{}, error) {
	// Парсинг JSON данных
	var jsonData map[string]interface{}
	err := json.Unmarshal(jsonFile, &jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func parseBanner(jsonData []byte) (Banner, error) {
	var banner Banner
	err := json.Unmarshal(jsonData, banner)
	//err := json.Unmarshal([]byte(jsonData), banner)
	if err != nil {
		return Banner{}, err
	}
	return banner, nil
}*/
