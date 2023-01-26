package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func createTableIfNotExists(db *sql.DB) {
	log.Print("Creating table if not exists")
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS heartbeats (id TEXT PRIMARY KEY, branch TEXT, category TEXT, created_at TEXT, cursorpos INTEGER, dependencies TEXT, entity TEXT, is_write INTEGER, language TEXT, lineno INTEGER, lines INTEGER, machine_name_id TEXT, project TEXT, project_root_count INTEGER, time REAL, type TEXT, user_agent_id TEXT, user_id TEXT)")
	if err != nil {
		panic(err)
	}
}

func arrayToString(array []string) string {
	var result string
	for _, item := range array {
		result += item + ","
	}
	return result
}

func insertHeartBeat(db *sql.DB, heartBeat HeartBeatResp) {
	log.Print("Inserting heartbeats into database")
	for _, heartBeat := range heartBeat.Data {
		_, err := db.Exec("INSERT OR REPLACE INTO heartbeats (id, branch, category, created_at, cursorpos, dependencies, entity, is_write, language, lineno, lines, machine_name_id, project, project_root_count, time, type, user_agent_id, user_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", heartBeat.Id, heartBeat.Branch, heartBeat.Category, heartBeat.CreatedAt, heartBeat.Cursorpos, arrayToString(heartBeat.Dependencies), heartBeat.Entity, heartBeat.IsWrite, heartBeat.Language, heartBeat.Lineno, heartBeat.Lines, heartBeat.MachineNameId, heartBeat.Project, heartBeat.ProjectRootCount, heartBeat.Time, heartBeat.Type, heartBeat.UserAgentId, heartBeat.UserId)
		if err != nil {
			panic(err)
		}
	}
}

func connectToDB(name string) *sql.DB {
	db, err := sql.Open("sqlite3", name)
	if err != nil {
		panic(err)
	}
	return db
}
