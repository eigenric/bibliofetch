package model

import (
	"errors"
	"reflect"
	"strconv"
)

type Libro struct{
	Edicion 		int `default:"1"`
	AnioPublicacion int `default:"1"`

	Editorial string
	DatosLibro DatosClave
}

// Comprueba que un dato sea positivo
func EsPositivo (num int) bool{
	return num >= 0
}

// Establece valores por defecto si no existen (función privada)
func setDefaults (obj interface{}) {
	v := reflect.ValueOf(obj).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		variable := v.Field(i)
		tipoVariable := t.Field(i)

		// Comprueba si es un entero y si tiene el valor por defecto (0)
		if variable.Kind() == reflect.Int && variable.Int() == 0 {

			// Comprueba si tiene etiqueta "default"
			if valorDefault, ok := tipoVariable.Tag.Lookup("default"); ok {

				// Cambia el valor si es entero
				if enteroDefault, err := strconv.Atoi(valorDefault); err == nil {
					variable.SetInt(int64(enteroDefault))
				}
			}
		}
	}
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

	// Se asignan los valores por defecto si fueran necesarios
	setDefaults(libro)

	return libro, nil
}
