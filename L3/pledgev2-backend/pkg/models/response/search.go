package response

import "pledgev2-backend/pkg/models"

type Search struct {
	Count int64         `json:"count"`
	Rows  []models.Pool `json:"rows"`
}
