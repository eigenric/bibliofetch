# Milestones - Bibliofetch

## [M0] Modelado del Problema
- **Objetivo**: Comprender en profundidad las dificultades de los estudiantes al acceder y manejar la bibliografía, definiendo personas y sus "user journeys" al gestionar la bibliografía.
- **Entregable**: Documento que describe los user-journeys e historias de usuario al tratar con la bibliografía.
- **Viabilidad**: Validación mediante la revisión del documento por pares.

## [M1] Extracción y Estructuración de la Bibliografía
- **Objetivo**: Desarrollar la funcionalidad para extraer datos bibliográficos de las guías docentes (PDF o web) y organizarlos de manera estructurada. 
- **Entregable**: Código del módulo que extrae la información bibliográfica en formato web y otro módulo para pdf y la organiza en un formato estructurado por atributos (título, autor, editorial, año, ISBN, etc.) 
- **Viabilidad**: Se verificará para un conjunto de guías de ejemplo mediante tests.

## [M2] Bibliografía según Grado, Curso e Idioma
- **Objetivo**: Simplificar la búsqueda de bibliografía permitiendo a los usuarios obtener la bibliografía según su grado y curso, filtrando por idioma. (Requiere superar el Milestone M1)
- **Entregable**: Módulos que, dado el grado y curso del usuario e idioma de preferencia, identifica las asignaturas asociadas al curso del grado y muestra la bibliografía correspondiente en ese idioma (mediante el módulo de extracción).
- **Viabilidad**: Validación probando el módulo con diferentes grados y cursos, verificando la precisión de la selección y presentación de la bibliografía.