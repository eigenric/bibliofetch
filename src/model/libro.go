package model

import (
	"errors"
)

type Libro struct{
	Edicion int 
	AnioPublicacion int 

	Editorial string
	DatosLibro DatosClave
}

func EsPositivo (valor int) bool{
	return valor >= 0
}

func NuevoLibro(edicion int, publicacion int, editorial string, datos DatosClave) (*Libro, error){
	edicionValida := EsPositivo(edicion)
	publicacionValida := EsPositivo(publicacion)

	if !edicionValida && !publicacionValida {
		return nil, errors.New("La edición y la fecha de publicación deben ser positivas")
	} else if !edicionValida {
		return nil, errors.New("El número edición debe ser positivo")
	} else if !publicacionValida { 
		return nil, errors.New("El año de publicación debe ser positivo")
	}

	libro := &Libro{edicion, publicacion, editorial, datos}

	return libro, nil
}
