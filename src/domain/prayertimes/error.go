package prayertimes

import "errors"

var (
	ErrInvalidRequest      = errors.New("invalid prayertimes request")
	ErrProviderUnavailable = errors.New("prayertimes provider unavailable")
	ErrExternalAPI         = errors.New("external API error")
	ErrParsingAPIResponse  = errors.New("failed to parse API response")
)
