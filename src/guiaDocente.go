package model

import (
	"errors"
)

type guiaDocente struct{
	anio int
	asignatura string
	
	bibliografiaBasica []libro			// Lista de libros básicos
	bibliografiaComplementaria []libro  // Lista de libros complementarios
}


// Crea una nueva guía docente con datos válidos
func nuevaGuiaDocente(anio int, asignatura string, basica []libro, complementaria []libro) (*guiaDocente, error){
	bool anioValido = esPositivo(anio)

	// Ambos datos inválidos
	if !anioValido {
		return nil, errors.New("El año debe ser positivo")
	}

	// Datos válidos: se crea la guía docente sin devolver errores
	return &guiaDocente{anio, asignatura, basica, complementaria}, nil
}
