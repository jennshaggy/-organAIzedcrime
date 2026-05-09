package models

import "encoding/json"
// ExternalReference holds the ATLAS ID (e.g. AML.TA0000) and source URL
type ExternalReference struct {
	SourceName string `json:"source_name"`
	ExternalID string `json:"external_id"`
	URL        string `json:"url"`
}

// KillChainPhase links a technique to its parent tactic
type KillChainPhase struct {
	KillChainName string `json:"kill_chain_name"`
	PhaseName     string `json:"phase_name"`
}

// Tactic maps to STIX objects of type "x-mitre-tactic"
type Tactic struct {
	Type               string              `json:"type"`
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	Description        string              `json:"description"`
	ExternalReferences []ExternalReference `json:"external_references"`
}

// Technique maps to STIX objects of type "attack-pattern"
type Technique struct {
	Type               string              `json:"type"`
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	Description        string              `json:"description"`
	KillChainPhases    []KillChainPhase    `json:"kill_chain_phases"`
	ExternalReferences []ExternalReference `json:"external_references"`
	IsSubtechnique     bool                `json:"x_mitre_is_subtechnique"`
}

// Bundle is the top-level STIX wrapper that contains all objects
type Bundle struct {
	Type    string        `json:"type"`
	Objects []json.RawMessage `json:"objects"`
}
