package models

import (
"encoding/json"
"fmt"
)

type CardAbility struct {
	Type   string  `json:"type"`
	Name   *string `json:"name,omitempty"`
	Effect *string `json:"effect,omitempty"`
}

type Damage string

func (d *Damage) UnmarshalJSON(data []byte) error {
	var i int
	if err := json.Unmarshal(data, &i); err == nil {
		*d = Damage(fmt.Sprintf("%d", i))
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		*d = Damage(s)
		return nil
	}
	return fmt.Errorf("invalid damage type")
}

type CardAttack struct {
	Name   *string  `json:"name,omitempty"`
	Cost   []string `json:"cost,omitempty"`
	Effect *string  `json:"effect,omitempty"`
	Damage *Damage  `json:"damage,omitempty"`
}

type CardItem struct {
	Name   *string `json:"name,omitempty"`
	Effect *string `json:"effect,omitempty"`
}

type CardVariants struct {
	Normal       bool `json:"normal"`
	Reverse      bool `json:"reverse"`
	Holo         bool `json:"holo"`
	FirstEdition bool `json:"firstEdition"`
	WPromo       bool `json:"wPromo"`
}

type CardWeakRes struct {
	Type  string  `json:"type"`
	Value *string `json:"value,omitempty"`
}

type Legal struct {
	Standard bool `json:"standard"`
	Expanded bool `json:"expanded"`
}

type SetCardCount struct {
	Total    int  `json:"total"`
	Official int  `json:"official"`
	Normal   *int `json:"normal,omitempty"`
	Reverse  *int `json:"reverse,omitempty"`
	Holo     *int `json:"holo,omitempty"`
	FirstEd  *int `json:"firstEd,omitempty"`
}

type Booster struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
