package command

import (
	"github.com/achmdndy/safa-life-api/src/domain/resource"
)

// CommandHandler handles all write operations for the Resource module.
type CommandHandler struct {
	resourceService resource.ResourceService
	txManager       resource.TransactionManager
}

// NewCommandHandler creates a new command handler.
func NewCommandHandler(service resource.ResourceService, txManager resource.TransactionManager) *CommandHandler {
	return &CommandHandler{
		resourceService: service,
		txManager:       txManager,
	}
}
