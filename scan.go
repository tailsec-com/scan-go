package scan

import (
	"encoding/json"
	"os/exec"
	"path/filepath"
)

type Client struct {
	Scanner string
	Binary  string
}

type ScanResult struct {
	Findings []Finding `json:"findings"`
	Summary  Summary  `json:"summary"`
}

type Finding struct {
	RuleID   string `json:"ruleId"`
	Severity string `json:"severity"`
	Title    string `json:"title"`
	Resource string `json:"resource"`
	File     string `json:"file,omitempty"`
	Line     int    `json:"line,omitempty"`
}

type Summary struct {
	Total int `json:"total"`
}

func NewClient(scanner string) *Client {
	return &Client{
		Scanner: scanner,
		Binary:  "npx",
	}
}

func (c *Client) Scan(path string) (*ScanResult, error) {
	args := []string{"@tailsec/scan-" + c.Scanner, path, "--format", "json"}
	cmd := exec.Command(c.Binary, args...)
	cmd.Dir = filepath.Dir(path)

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var result ScanResult
	if err := json.Unmarshal(output, &result); err != nil {
		return nil, err
	}

	return &result, nil
}