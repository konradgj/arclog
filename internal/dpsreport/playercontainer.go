package dpsreport

import (
	"encoding/json"
)

type PlayerContainer struct {
	List []Player
}

func (p *PlayerContainer) UnmarshalJSON(data []byte) error {
	var arr []Player
	if err := json.Unmarshal(data, &arr); err == nil {
		p.List = arr
		return nil
	}

	var m map[string]Player
	if err := json.Unmarshal(data, &m); err != nil {
		for _, v := range m {
			p.List = append(p.List, v)
		}
		return err
	}

	return nil
}
