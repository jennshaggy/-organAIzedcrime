package loader

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jennshaggy/organAIzedcrime/models"
)

func ParseBundle(path string) ([]models.Tactic, []models.Technique, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, fmt.Errorf("could not read file: %w", err)
	}

	var bundle models.Bundle
	if err := json.Unmarshal(data, &bundle); err != nil {
		return nil, nil, fmt.Errorf("could not parse bundle: %w", err)
	}

	var tactics []models.Tactic
	var techniques []models.Technique

	for _, raw := range bundle.Objects {
		var typed struct {
			Type string `json:"type"`
		}
		if err := json.Unmarshal(raw, &typed); err != nil {
			continue
		}

		switch typed.Type {
		case "x-mitre-tactic":
			var t models.Tactic
			if err := json.Unmarshal(raw, &t); err == nil {
				tactics = append(tactics, t)
			}
		case "attack-pattern":
			var t models.Technique
			if err := json.Unmarshal(raw, &t); err == nil {
				techniques = append(techniques, t)
			}
		}
	}

	return tactics, techniques, nil
}