package spec

const DB_ADMIN_SUB = "tasks.db_admin"
const CREATE_DB_SUB = DB_ADMIN_SUB + ".create_db.%s"

type DeleteDB struct {
	ContainerName string
}

type CreateDB struct {
	Image    string
	Name     string
	Password string

	// Path Optional filesystem path for local/embedded engines
	Path string
	Port string

	// Role Database role for connection
	Role     string
	SslMode  string
	Username string
}

type Status string

const (
	StatusPass Status = "PASS"
	StatusFail Status = "FAIL"
)

type AgentResponse struct {
	Message string
	Status  Status // "PASS or FAIL"
}
