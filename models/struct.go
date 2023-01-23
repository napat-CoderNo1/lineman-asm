package models

import "encoding/json"

type DataStruct struct {
	Data [] Data `json:"Data"`
}

type Data struct {
	ConfirmDate    string          `json:"ConfirmDate"`
	No             string          `json:"No"`
	Age            json.RawMessage `json:"Age"`
	Gender         string          `json:"Gender"`
	GenderEn       string          `json:"GenderEn"`
	Nation         string          `json:"Nation"`
	NationEn       string          `json:"NationEn"`
	Province       string          `json:"Province"`
	ProvinceId     int             `json:"ProvinceId"`
	District       string          `json:"District"`
	ProvinceEn     string          `json:"ProvinceEn"`
	StatQuarantine int             `json:"StatQuarantine"`
}

type Response struct {
	Province map[string]int `json:"Province"`
	AgeGroup map[string]int `json:"AgeGroup"`
}