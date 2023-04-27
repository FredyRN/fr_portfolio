package models

import (
	"github.com/google/uuid"
)

type Country struct {
	ID          uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	CountryName string    `json:"País"`
}

type State struct {
	ID        uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	StateName string    `json:"Departamento"`
	Country   uuid.UUID `bun:"type:uuid"`
}

type City struct {
	ID       uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	CityName string    `json:"Ciudad"`
	State    uuid.UUID `bun:"type:uuid"`
}

type Address struct {
	ID           uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	Street       string    `json:"Calle"`
	City         uuid.UUID `bun:"type:uuid"`
	Address      string    `json:"Dirección"`
	Neighborhood string    `json:"Barrio"`
	PostalCode   string    `json:"Código postal"`
}

type Profile struct {
	ID             uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	FirstName      string    `json:"Nombre"`
	LastName       string    `json:"Apellido"`
	DocumentType   string    `json:"Tipo de documento"`
	Identification string    `json:"Número de identificación"`
	Address        uuid.UUID `bun:"type:uuid"`
	EmailAddress   string    `json:"Email"`
	PhoneNumber    string    `json:"Teléfono"`
	BirthDate      string    `json:"Fecha de nacimiento"`
	LastModified   string    `json:"Última actualización"`
}

type User struct {
	ID       uuid.UUID `bun:"type:uuid"`
	Username string    `json:"Usuario"`
	Password string    `json:"Contraseña"`
}
