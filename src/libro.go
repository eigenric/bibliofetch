package model

import (
	"errors"
)

type libro struct{
	edicion int
	anioPublicacion int

	editorial string

	datosClave datosClave // Estructura de variables de identificación unívoca de un libro
}

// Comprueba que un dato sea positivo
func esPositivo (num int) bool{
	return num > 0
}

// Crea un nuevo libro con datos válidos
func nuevoLibro(edicion int, publicacion int, editorial string, datos datosClave) (*libro, error){
	edicionValida := esPositivo(edicion)
	publicacionValida := esPositivo(publicacion)

	// Ambos datos inválidos
	if !edicionValida && !publicacionValida {
		return nil, errors.New("La edición y la fecha de publicación deben ser positivas")
	} else if !edicionValida { // Edición inválida
		return nil, errors.New("El número edición debe ser positivo")
	} else if !publicacionValida { // Año de publicación inválida
		return nil, errors.New("El año de publicación debe ser positivo")
	}

	// Datos válidos: se crea el libro sin devolver errores
	return &libro{edicion, publicacion, editorial, datos}, nil
}
