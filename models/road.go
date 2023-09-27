package models

import "gorm.io/gorm"

type Road_logs struct {
	gorm.Model
	Latitud   string `json:"latitud"`
	Longitud  string `json:"longitud"`
	Fecha     string `json:"fecha"`
	Velocidad string `json:"velocidad"`
	Alerta    string `json:"alerta"`
	UserID    uint   `json:"userid"`
}
