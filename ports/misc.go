package ports

type MiscRepository interface {
	GetDatabaseStatus() bool
}
