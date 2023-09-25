package models

import "gorm.io/gorm"

type Userdata struct {
	gorm.Model
	Correo              string      `json:"Correo" validate:"required" gorm:"unique"`
	Password            string      `json:"Password" validate:"required"`
	Nombre              string      `json:"Nombre" validate:"required"`
	Apellido            string      `json:"Apellido" validate:"required"`
	Fecha_nacimiento    string      `json:"fecha_nacimiento" validate:"required"`
	Numero_telefono     string      `json:"numero_telefono" validate:"required"`
	Contacto_emergencia string      `json:"contacto_emergencia" validate:"required"`
	Road_logs           []Road_logs `json:"road_logs"`
}
