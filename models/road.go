package models

import "gorm.io/gorm"

type Road_logs struct {
	gorm.Model
	Latitud   string
	Longitud  string
	Fecha     string
	Velocidad string
	Alerta    string
}
