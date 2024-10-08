# Milestones - Bibliofetch

## [M0] Modelado del Problema
- **Objetivo**: Comprender en profundidad las dificultades de los estudiantes al acceder y manejar la bibliografía, definiendo personas y sus "user journeys" al gestionar la bibliografía.
- **Entregable**: Documento que describe los user-journeys e historias de usuario al tratar con la bibliografía.
- **Viabilidad**: Validación mediante la revisión del documento por pares.

## [M1] Bibliofetch extrae la bibliografía de las guías docentes
- **Objetivo**: Desarrollar la funcionalidad para extraer datos bibliográficos de las guías docentes (PDF o web) y organizarlos de manera estructurada. 
- **Entregable**: Módulo que recibe una guía docente como entrada (en formato pdf o web), extrae la información bibliográfica y la organiza en un formato estructurado por atributos (título, autor, editorial, año, ISBN, etc.). 
- **Viabilidad**: Se etiquetarán manualmente algunas guías docentes y se testeará el módulo para verificar la corrección de la extracción.

## [M2] Bibliofetch filtra la bibliografía según Idioma.
- **Objetivo**: Reconocer el idioma de la bibliografía y filtrarla según el idioma preferido por el usuario.
(Requiere superar el Milestone M1)
- **Entregable**: Módulo que, dado un conjunto de libros, identifica el idioma de cada libro y filtra los libros según el idioma preferido por el usuario.
- **Viabilidad**: Se etiquetarán manualmente algunos libros con su idioma y se testeará el módulo para verificar la corrección del filtrado.

## [M3] Bibliofetch devuelve las guías docentes de las asignaturas según Grado y Curso
- **Objetivo**: Simplificar la búsqueda de bibliografía permitiendo a los usuarios obtener la bibliografía según su grado, curso e idioma. (Requiere superar el Milestone M1 y M2 )
- **Entregable**: Módulos que, dado el grado y curso del usuario e idioma de preferencia, identifica las asignaturas asociadas al curso del grado y muestra la bibliografía correspondiente en ese idioma (usando el módulo de extracción y el de filtrado por idioma).
- **Viabilidad**: Se obtendrá un conjunto de datos que relaciona grados, cursos, asignaturas y guías docentes, y se testeará el módulo para verificar la corrección de la búsqueda.