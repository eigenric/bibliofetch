# Milestones - Bibliofetch

## [M0] Especificación del dominio del problema
**Historias de Usuario**: [HU1]

**Objetivo**: A partir de las Historias de Usuario (HUs), reconocer los elementos clave del dominio y sus relaciones con el fin de modelar un sistema que facilite el acceso y gestión de la bibliografía de las asignaturas de los estudiantes.

**Entregables:**
    - Código que represente los elementos del dominio, con estructuras para Libro, Asignatura, Bibliografía.
    - De igual manera, las relaciones entre los elementos debe ser modelada, con atención a los atributos del Libro.

**Viabilidad:** 
    - El modelo es válido si incorpora todas las entidades esenciales del dominio y sus relaciones, correctamente implementadas.
    - Se considerará viable si la formulación del modelo de asignatura permite priorizar unos libros frente a otros.
    - Se considera viable cuando otro desarrollador puede comprender el modelo y continuar su desarrollo a partir de la especificación.    

## [M1] Bibliofetch extrae la bibliografía de las guías docentes

**Historias de Usuario**: [HU1]
- **Objetivo**: Desarrollar la funcionalidad para extraer datos bibliográficos de las guías docentes (PDF o web) y organizarlos de manera estructurada. 
- **Entregable**: Módulo que recibe una guía docente como entrada web, extrae la información bibliográfica y la organiza en un formato estructurado por atributos (título, autor, editorial, año, ISBN, etc.). Requiere [M0].
- **Viabilidad**: Se etiquetarán manualmente algunas guías docentes y se testeará el módulo para verificar la corrección de la extracción.

## [M2] Bibliofetch permite seleccionar libros favoritos

**Historias de Usuario**: [HU2]

- **Objetivo**: Desarrollar la funcionalidad para que los estudiantes puedan seleccionar libros favoritos de la bibliografía extraída, con el fin de priorizar su estudio.
- **Entregable**: Módulo que permite al usuario seleccionar sus libros favoritos, a partir de los datos bibliográficos extraídos. Requiere [M1].
- **Viabilidad**: Se testeará el módulo para verificar la correcta selección de libros favoritos.

## [M3] Bibliofetch exporta la bibliografía en formato BibTeX
**Historias de Usuario**: [HU4]

**Objetivo**: A partir de los datos bibliográficos extraídos, generar un archivo BibTeX que permita a los usuarios citarla en sus trabajos académicos.
**Entregable**: Módulo que recibe los datos bibliográficos estructurados y genera un archivo BibTeX con los atributos de los libros. Requiere [M1].
**Viabilidad**: Se testeará el módulo para verificar la corrección de la exportación del archivo BibTeX.