package command

import (
	"github.com/achmdndy/safa-life-api/src/application/resource/dto"
)

func FromCreateResourceRequest(req dto.CreateResourceRequest) CreateResourceCommand {
	return CreateResourceCommand{
		ID:            req.ID,
		Type:          req.Type,
		Name:          req.Name,
		Language:      req.Language,
		Version:       req.Version,
		LastUpdatedAt: req.LastUpdatedAt,
		DownloadURL:   req.DownloadURL,
		Size:          req.Size,
	}
}

func FromUpdateResourceRequest(req dto.UpdateResourceRequest) UpdateResourceCommand {
	cmd := UpdateResourceCommand{ID: req.ID}
	if req.Type != "" {
		cmd.Type = &req.Type
	}
	if req.Name != "" {
		cmd.Name = &req.Name
	}
	if req.Language != "" {
		cmd.Language = &req.Language
	}
	if req.Version != "" {
		cmd.Version = &req.Version
	}
	if !req.LastUpdatedAt.IsZero() {
		cmd.LastUpdatedAt = &req.LastUpdatedAt
	}
	if req.DownloadURL != "" {
		cmd.DownloadURL = &req.DownloadURL
	}
	if req.Size > 0 {
		cmd.Size = &req.Size
	}
	return cmd
}

func FromUpdateInstallStatusRequest(req dto.UpdateInstallStatusRequest) UpdateInstallStatusCommand {
	return UpdateInstallStatusCommand{
		ID:          req.ID,
		IsInstalled: req.IsInstalled,
	}
}
