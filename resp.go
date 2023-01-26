package main

import "time"

type HeartBeatResp struct {
	Data []struct {
		Branch           string    `json:"branch"`
		Category         string    `json:"category"`
		CreatedAt        time.Time `json:"created_at"`
		Cursorpos        int       `json:"cursorpos"`
		Dependencies     []string  `json:"dependencies"`
		Entity           string    `json:"entity"`
		Id               string    `json:"id"`
		IsWrite          bool      `json:"is_write"`
		Language         string    `json:"language"`
		Lineno           int       `json:"lineno"`
		Lines            int       `json:"lines"`
		MachineNameId    string    `json:"machine_name_id"`
		Project          string    `json:"project"`
		ProjectRootCount int       `json:"project_root_count"`
		Time             float64   `json:"time"`
		Type             string    `json:"type"`
		UserAgentId      string    `json:"user_agent_id"`
		UserId           string    `json:"user_id"`
	} `json:"data"`
	End      time.Time `json:"end"`
	Start    time.Time `json:"start"`
	Timezone string    `json:"timezone"`
}
