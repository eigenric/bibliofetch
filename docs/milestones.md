# Milestones - Bibliofetch

Se ha planificado el proyecto de forma incremental en una serie de tres milestones ([M0],[M1],[M2])
Cada milestone representa un PMV (Producto Mínimamente Viable).

## [M0] Especificación del problema

Partir de las historias de usuario para reconocer los elementos clave del dominio y las relaciones entre ellos.
Elegir el lenguaje de programación con el que se va a resolver el problema y ser consciente de sus buenas prácticas.

**Entregables:** Código con las estructuras fundamentales relativas a la Gestión Bibliográfica y sus relaciones

**Viabilidad:** 
    - Es viable si la formulación del modelo de asignatura permite distinguir libros prioritarios.
    - Se considera viable cuando otro desarrollador puede comprender el modelo y continuar su desarrollo a partir de la especificación.    

## [M1] Bibliofetch extrae la bibliografía de las guías docentes

- **Mínimo**: Bibliofetch puede extraer datos bibliográficos de las guías docentes (PDF o web) y organizarlos de manera estructurada. 
- **Entregable**: Módulo que recibe una guía docente como entrada web, extrae la información bibliográfica y la organiza en un formato estructurado por atributos (título, autor, editorial, año, ISBN, etc.). Requiere [M0].
- **Viabilidad**: Se etiquetarán manualmente algunas guías docentes y se testeará el módulo para verificar la corrección de la extracción.

## [M2] Bibliofetch permite seleccionar libros favoritos

- **Mínimo**: El producto permite que los estudiantes puedan seleccionar libros favoritos de la bibliografía extraída, con el fin de priorizar su estudio.
- **Entregable**: Módulo que permite al usuario seleccionar sus libros favoritos, a partir de los datos bibliográficos extraídos. Requiere [M1].
- **Viabilidad**: Se testeará el módulo para verificar la correcta selección de libros favoritos.

## [M3] Bibliofetch exporta la bibliografía en formato BibTeX

**Mínimo**: A partir de los datos bibliográficos extraídos, el producto permite generar un archivo BibTeX listo para citar en trabajos académicos.
**Entregable**: Módulo que recibe los datos bibliográficos estructurados y genera un archivo BibTeX con los atributos de los libros. Requiere [M1].
**Viabilidad**: Se testeará el módulo para verificar la corrección del formato del archivo (BibTeX).