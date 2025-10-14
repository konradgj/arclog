package dpsreport

type UploadResponse struct {
	ID               string     `json:"id"`
	Permalink        string     `json:"permalink"`
	UploadTime       int64      `json:"uploadTime"`
	EncounterTime    int64      `json:"encounterTime"`
	Generator        string     `json:"generator"`
	GeneratorID      int        `json:"generatorId"`
	GeneratorVersion int        `json:"generatorVersion"`
	Language         string     `json:"language"`
	LanguageID       int        `json:"languageId"`
	Evtc             EvtcMeta   `json:"evtc"`
	Players          []Player   `json:"players"`
	Encounter        Encounter  `json:"encounter"`
	Report           ReportInfo `json:"report"`
	Error            *string    `json:"error"`
	UserToken        string     `json:"userToken"`
}

type Encounter struct {
	UniqueID        string  `json:"uniqueId"`
	Success         bool    `json:"success"`
	Duration        float64 `json:"duration"`
	CompDps         int     `json:"compDps"`
	NumberOfPlayers int     `json:"numberOfPlayers"`
	NumberOfGroups  int     `json:"numberOfGroups"`
	BossID          int     `json:"bossId"`
	Boss            string  `json:"boss"`
	IsCm            bool    `json:"isCm"`
	Gw2Build        int     `json:"gw2Build"`
	JSONAvailable   bool    `json:"jsonAvailable"`
}

type EvtcMeta struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	BossID  int    `json:"bossId"`
}

type Player struct {
	DisplayName   string `json:"display_name"`
	CharacterName string `json:"character_name"`
	Profession    int    `json:"profession"`
	EliteSpec     int    `json:"elite_spec"`
}

type ReportInfo struct {
	Anonymous bool `json:"anonymous"`
	Detailed  bool `json:"detailed"`
}
