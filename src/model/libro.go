package model

import (
	"errors"
	"reflect"
	"strconv"
)

type Libro struct{
	Edicion int 
	AnioPublicacion int 

	Editorial string
	DatosLibro DatosClave
}

// Comprueba que un dato sea positivo
func EsPositivo (valor int) bool{
	return valor >= 0
}

// Crea un nuevo libro con datos válidos
func NuevoLibro(edicion int, publicacion int, editorial string, datos DatosClave) (*Libro, error){
	edicionValida := EsPositivo(edicion)
	publicacionValida := EsPositivo(publicacion)

	// Ambos datos inválidos
	if !edicionValida && !publicacionValida {
		return nil, errors.New("La edición y la fecha de publicación deben ser positivas")
	} else if !edicionValida { // Edición inválida
		return nil, errors.New("El número edición debe ser positivo")
	} else if !publicacionValida { // Año de publicación inválida
		return nil, errors.New("El año de publicación debe ser positivo")
	}

	// Datos válidos: se crea el libro sin devolver errores
	libro := &Libro{edicion, publicacion, editorial, datos}

	return libro, nil
}
