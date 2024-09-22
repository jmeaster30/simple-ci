package main

import (
	"time"

	"github.com/jmeaster30/twilite"
)

type Action struct {
	RepositoryUrl         string
	Action                WorkflowAction
	If                    string
	BuildCommand          string
	ArtifactUploadCommand string
}

type Job struct {
	Action    Action
	StartTime time.Time
	EndTime   time.Time
	HookId    string
	JobState  JobState
	ProcessId uint
	Output    string
}

type JobState uint8

const (
	JobStateCreated JobState = iota
	JobStateRunning
	JobStateSuccess
	JobStateFailed
	JobStateCanceled
)

type Database struct {
	context *twilite.DbContext
}

var databaseInstance *Database

func GetDatabase(databaseFilename string) *Database {
	if databaseInstance != nil {
		return databaseInstance
	}

	context := twilite.NewDbContext(databaseFilename)

	twilite.RegisterTable[Action](&context)
	twilite.RegisterTable[Job](&context)

	databaseInstance = &Database{
		context: &context,
	}

	return databaseInstance
}

func (db *Database) Action() twilite.QueryBuilder[Action] {
	return twilite.SelectQuery[Action](db.context)
}

func (db *Database) Job() twilite.QueryBuilder[Job] {
	return twilite.SelectQuery[Job](db.context)
}
