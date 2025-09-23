package health

import (
	"fmt"

	"github.com/achmdndy/safa-life-api/src/domain/health"
	"github.com/achmdndy/safa-life-api/src/infrastructure/configs"
)

type CheckerRepository struct {
	db    configs.DatabaseInterface
	redis configs.RedisInterface
}

func NewCheckerRepository(db configs.DatabaseInterface, redis configs.RedisInterface) health.CheckerRepository {
	return &CheckerRepository{
		db:    db,
		redis: redis,
	}
}

func (c *CheckerRepository) CheckDatabase() error {
	if c.db == nil {
		return fmt.Errorf("database connection is nil")
	}
	return c.db.Ping()
}

func (c *CheckerRepository) CheckRedis() error {
	if c.redis == nil {
		return fmt.Errorf("redis connection is nil")
	}
	return c.redis.Ping()
}