package handler

import "go-Framework/app/services"

type Handler struct {
	Session *Session
}

func NewHandler(svc *services.Services) *Handler {
	return &Handler{
		Session: NewSession(svc.Session),
	}
}
