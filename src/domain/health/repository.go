package health

type CheckerRepository interface {
	CheckDatabase() error
	CheckRedis() error
}
