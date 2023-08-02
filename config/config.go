package config

import (
	"fmt"
	"os"

	"github.com/kardianos/osext"
)

var (
	AppName           = "agenda"
	AppVersion        = GetEnvOrDefault("AGENDA_APP_VERSION", "1.0.0")
	RestPort          = GetEnvOrDefault("AGENDA_REST_PORT", "12580")
	GrpcPort          = GetEnvOrDefault("AGENDA_GRPC_PORT", "12581")
	PidPath           = GetEnvOrDefault("AGENDA_PID_PATH", fmt.Sprintf("/tmp/%s.pid", AppName))
	Env               = GetEnvOrDefault("AGENDA_ENV", "dev")
	LogPath           = GetEnvOrDefault("AGENDA_LOG_PATH", fmt.Sprintf("/tmp/%s.out", AppName))
	EnableMetrics     = GetEnvOrDefaultBool("AGENDA_ENABLE_METRICS", true)
	UseMigration      = GetEnvOrDefaultBool("AGENDA_MIGRATION_ENABLED", true)
	MigrationPath     = GetEnvOrDefault("AGENDA_MIGRATION_PATH", "")
	MigrationSkipInit = GetEnvOrDefaultBool("AGENDA_MIGRATION_SKIP_INIT", false)
	EnableAuditing    = GetEnvOrDefaultBool("AGENDA_ENABLE_AUDITING", false)
	AuditStorage      = GetEnvOrDefault("AGENDA_AUDIT_STORAGE", "file") // file or sql
	AuditFilePath     = GetEnvOrDefault("AGENDA_AUDIT_FILE_PATH", "./audit")
	LogOutput         = os.Stdout
	ExecDir           = ""
	StorageSQL        = true
	SqlitePath        = GetEnvOrDefault("AGENDA_SQLITE_PATH", "")
	DocumentDir       = GetEnvOrDefault("AGENDA_DOCUMENT_DIR", "./documents")
)

func init() {
	ExecDir, _ = osext.ExecutableFolder()
}

func GetEnvOrDefault(envVar string, defEnvVar string) (newEnvVar string) {
	if newEnvVar = os.Getenv(envVar); len(newEnvVar) == 0 {
		return defEnvVar
	} else {
		return newEnvVar
	}
}

func GetEnvOrDefaultBool(envVar string, defEnvVar bool) (newEnvVar bool) {
	newEnvVarStr := os.Getenv(envVar)
	if len(newEnvVarStr) == 0 {
		return defEnvVar
	}
	return newEnvVarStr == "true"
}

func GetUniqueNumberForLock() (number int) {
	for i := range AppName {
		number += int(AppName[i])
	}
	return
}
