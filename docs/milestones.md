# Milestones - Bibliofetch

Se ha planificado el proyecto de forma incremental en una serie de dos milestones ([M0],[M1])

Cada milestone representa un PMV (Producto Mínimamente Viable).

## [M0] Representación del problema

Definir las estructuras para representar el problema descrito en la [HU1](https://github.com/eigenric/bibliofetch/issues/2).
Elegir el lenguaje de programación con el que se va a resolver el problema.

**Entregables:** Código con las entidades fundamentales del problema.

**Viabilidad:** 
    - Se ha definido una estructura (clase, struct) para cada entidad que representa adecuadamente el problema.
    - Se utilizan las buenas prácticas del lenguaje de programación en la organización de archivos y entidades.

## [M1] Bibliofetch extrae la bibliografía de las guías docentes

- **Mínimo**: Bibliofetch puede extraer datos bibliográficos de las guías docentes (PDF o web) y organizarlos de manera estructurada. Permitiría solventar la necesidad planteada en la [HU1](https://github.com/eigenric/bibliofetch/issues/2).
- **Entregable**: Módulo que recibe una guía docente como entrada web, extrae la información bibliográfica y la organiza en un formato estructurado por atributos de acuerdo al modelo de representación y en particular a la necesidad del usuario [HU1]. Requiere [M0].
- **Viabilidad**: Se etiquetarán manualmente algunas guías docentes y se testeará el módulo para verificar la corrección de la extracción.