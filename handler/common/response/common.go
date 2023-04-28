package common

const (
	Success = 0
	Failed  = 1

	// Base 3000 is Custom Error Code JON
	ErrorFailedToBindJSON = 3000

	// Base 4000 is Custom Error Code
	ErrorDatabaseFailure =  4000
	ErrorDatabaseFailedStatement = 4001
	ErrorDatabaseQueryExecution  = 4002
	ErrorDatabaseInsertion       = 4003
	ErrorDatabaseUpdate      = 4004
	ErrorDatabaseDelete      = 4005
	ErrorDatabaseFind        = 4006
)
