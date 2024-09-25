package app

import "time"

type Dash struct {
	DashId    string    `json:"dashId"`
	Name      string    `json:"name"`
	Speed     int       `json:"speed"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateDashRequest struct {
	Name  string `json:"name" binding:"required,min=1,max=10"`
	Speed int    `json:"speed" binding:"required"`
}

type CreateDashResponse struct {
	Dash Dash `json:"dash"`
}

type ListDashRequest struct {
}

type ListDashResponse struct {
	Dashes []Dash `json:"dashes"`
}
