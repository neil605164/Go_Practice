package handler

import "Go_Practice/app/business"

type Handler struct {
	// 繼承 business interface
	BInter business.BInter
}

func NewHandler() *Handler {
	return &Handler{
		BInter: business.NewBusiness(),
	}
}
