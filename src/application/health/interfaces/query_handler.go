package interfaces

// QueryHandler defines the interface for health query handlers in application layer
// This interface is used by presentation layer to avoid importing domain layer
type QueryHandler interface {
	Handle(query interface{}) (interface{}, error)
}