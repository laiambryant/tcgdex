package models

type Card struct {
	CardResume
	Illustrator    *string       `json:"illustrator,omitempty"`
	Rarity         string        `json:"rarity"`
	Category       string        `json:"category"`
	Variants       CardVariants  `json:"variants"`
	Set            SetResume     `json:"set"`
	DexID          []int         `json:"dexId,omitempty"`
	HP             *int          `json:"hp,omitempty"`
	Types          []string      `json:"types,omitempty"`
	EvolveFrom     *string       `json:"evolveFrom,omitempty"`
	Description    *string       `json:"description,omitempty"`
	Level          *string       `json:"level,omitempty"`
	Stage          *string       `json:"stage,omitempty"`
	Suffix         *string       `json:"suffix,omitempty"`
	Item           *CardItem     `json:"item,omitempty"`
	Abilities      []CardAbility `json:"abilities,omitempty"`
	Attacks        []CardAttack  `json:"attacks,omitempty"`
	Weaknesses     []CardWeakRes `json:"weaknesses,omitempty"`
	Resistances    []CardWeakRes `json:"resistances,omitempty"`
	Retreat        *int          `json:"retreat,omitempty"`
	RegulationMark *string       `json:"regulationMark,omitempty"`
	Legal          Legal         `json:"legal"`
	Boosters       []Booster     `json:"boosters,omitempty"`
}
