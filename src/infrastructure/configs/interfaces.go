package configs

// DatabaseInterface defines the contract for database operations
type DatabaseInterface interface {
	Ping() error
	Close() error
}

// RedisInterface defines the contract for Redis operations
type RedisInterface interface {
	Ping() error
	Close() error
}