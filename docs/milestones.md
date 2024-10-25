# Milestones - Bibliofetch

Se ha planificado el proyecto de forma incremental en una serie de cuatro milestones ([M0],[M1],[M2],[M3]).

Cada milestone representa un PMV (Producto Mínimamente Viable).

## [M0] Representación del problema

En lo relativo a la primera historia de usuario, definir las estructuras que modelan el problema de Extracción de Bibliografía.
Elegir el lenguaje de programación con el que se va a resolver el problema.

**Entregables:** Código con las entidades fundamentales relativas a la Gestión Bibliográfica.

**Viabilidad:** 
    - Se utilizan las buenas prácticas del lenguaje de programación en la organización de archivos y entidades.

## [M1] Bibliofetch extrae la bibliografía de las guías docentes

- **Mínimo**: Bibliofetch puede extraer datos bibliográficos de las guías docentes (PDF o web) y organizarlos de manera estructurada. Permitiría solventar la necesidad planteada en la [HU1].
- **Entregable**: Módulo que recibe una guía docente como entrada web, extrae la información bibliográfica y la organiza en un formato estructurado por atributos (título, autor, editorial, año, ISBN, etc.). Requiere [M0].
- **Viabilidad**: Se etiquetarán manualmente algunas guías docentes y se testeará el módulo para verificar la corrección de la extracción.

## [M2] Bibliofetch permite seleccionar libros favoritos

- **Mínimo**: El producto permite que los estudiantes puedan seleccionar libros favoritos de la bibliografía extraída, con el fin de priorizar su estudio. Resolvería el problema descrito en la [HU2]
- **Entregable**: Módulo que permite al usuario seleccionar sus libros favoritos, a partir de los datos bibliográficos extraídos. Requiere [M1].
- **Viabilidad**: Se testeará el módulo para verificar la correcta selección de libros favoritos.

## [M3] Bibliofetch exporta la bibliografía en formato BibTeX

**Mínimo**: A partir de los datos bibliográficos extraídos, el producto permite generar un archivo BibTeX listo para citar en trabajos académicos. Completar este hito resolvería la [HU3], para estudiantes con TFG.
**Entregable**: Módulo que recibe los datos bibliográficos estructurados y genera un archivo BibTeX con los atributos de los libros. Requiere [M1].
**Viabilidad**: Se testeará el módulo para verificar la corrección del formato del archivo (BibTeX).