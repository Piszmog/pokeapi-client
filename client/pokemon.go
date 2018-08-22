package client

import "encoding/json"

type Pokemon struct {
	Id                            int         `json:"id"`
	Name                          string      `json:"name"`
	Order                         int         `json:"order"`
	BaseExperience                int         `json:"base_experience"`
	IsDefault                     bool        `json:"is_default"`
	Height                        int         `json:"height"`
	Weight                        int         `json:"weight"`
	Abilities                     []Ability   `json:"abilities"`
	Forms                         []Details   `json:"forms"`
	GameIndices                   []GameIndex `json:"game_indices"`
	Stats                         []Stats     `json:"stats"`
	Moves                         []Move      `json:"moves"`
	SpriteUrls                    SpriteUrls  `json:"sprites"`
	HeldItems                     []HeldItem  `json:"held_items"`
	LocationAreaEncountersUrlPath string      `json:"location_area_encounters"`
	Species                       Details     `json:"species"`
	Types                         []Type      `json:"types"`
}

type Ability struct {
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
	Ability  Details `json:"ability"`
}

type Details struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type GameIndex struct {
	GameIndex int     `json:"game_index"`
	Version   Details `json:"version"`
}

type Stats struct {
	Effort   int     `json:"effort"`
	BaseStat int     `json:"base_stat"`
	Stat     Details `json:"stat"`
}

type Move struct {
	Move                Details              `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

type VersionGroupDetail struct {
	LevelLearnedAt  int     `json:"level_learned_at"`
	MoveLearnMethod Details `json:"move_learn_method"`
	VersionGroup    Details `json:"version_group"`
}

type HeldItem struct {
	Version Details `json:"version"`
	Rarity  string  `json:"rarity"`
}

type SpriteUrls struct {
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
	BackDefault      string `json:"back_default"`
	FrontFemale      string `json:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female"`
	BackShiny        string `json:"back_shiny"`
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
}

type Type struct {
	Slot int     `json:"slot"`
	Type Details `json:"type"`
}

func (pokemon Pokemon) MarshalBinary() ([]byte, error) {
	return json.Marshal(pokemon)
}

func (pokemon Pokemon) UnmarshalBinary(bytes []byte) error {
	return json.Unmarshal(bytes, pokemon)
}

type Client interface {
	GetPokemon(identifier string) (*Pokemon, error)
}
