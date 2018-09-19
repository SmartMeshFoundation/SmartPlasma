package handlers

import (
	"context"
	"time"

	"github.com/SmartMeshFoundation/SmartPlasma/service"
)

// SmartPlasma implements PlasmaCash methods to RPC server.
type SmartPlasma struct {
	timeout int
	service *service.Service
}

// NewSmartPlasma creates new SmartPlasma handler service.
func NewSmartPlasma(timeout int, service *service.Service) *SmartPlasma {
	return &SmartPlasma{
		timeout: timeout,
		service: service,
	}
}

func (api *SmartPlasma) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(
		context.Background(), time.Duration(api.timeout)*time.Second)
}
