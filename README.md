# Bibliofetch

## Formulación del problema

**Cliente:** Estudiante de Ingeniería, Ciencias o Humanidades de la Universidad de Granada.
**Contexto:**  Con el objetivo de superar con éxito las asignaturas de su grado, los estudiantes deben consultar la bibliografía recomendada. Esta se divide en fundamental y complementaria y está disponible en un apartado de la guía docente de cada asignatura (en PDF y en versión web). 

Dado que quieren minimizar su tiempo y coste empleado en lo relativo a los libros, los estudiantes se encuentran con los siguientes problemas:

*  **Extracción de atributos**: La bibliografía recomendada se presenta en formato de texto plano, en lugar de tener un formato estructurado para los atributos de los libros como título, autor, editorial, año de publicación, ISBN, etc.
*   **Verificación de disponibilidad:**  La bibliografía no siempre está disponible en la biblioteca de la universidad o, si está, no se encuentra en forma de enlace directo al catálogo de la biblioteca. Si está en forma de enlace, puede estar caído.
*  **Distinción de la bibliografía esencial:** El exceso de bibliografía (incluso fundamental) supone un problema para la gestión del tiempo de estudio del cliente o incluso un problema económico si es necesario comprar algunos libros.
*  **Adecuación al idioma del estudiante:** Si la bibliografía recomendada está en un sólo idioma (inglés o español) puede ser un hándicap para estudiantes nativos en otro idioma.

**Ejemplo**: Guías Docentes de Informática

En la [Guía Docente](https://www.ugr.es/estudiantes/grados/grado-ingenieria-informatica/sistemas-operativos/guia-docente) de la asignatura de Sistemas Operativos del Grado en Informática, notamos como:

-  En algunos libros se utiliza la notación "3/e" para indicar la tercera edición de un libro mientras que en otro caso, se usa la notación (2 ed.). En otras guías docentes, se escribe directamente "3ª edición" o en inglés 2nd edition.
- Aparece un libro repetido (M. Kerrisk, The Linux Programming Interface, No Starch Press , 2010) en la bibliografía fundamental y en la complementaria. En un caso, con el autor M. Kerrisk (nombre correcto), y en otro con el autor Michael Kerrish. (¿Errata?).
- En ningún caso, se enlaza al catálogo de la biblioteca, debiendo de buscar el libro manualmente en Gratensis para verificar su disponibilidad.

En la [Guía Docente](https://www.ugr.es/estudiantes/grados/grado-ingenieria-informatica/computacubicua-inteligambiental-etecnolinf/guia-docente) de la asignatura de Computación Ubicua e Inteligencia Ambiental del Grado en Informática, notamos como:

- Toda la bibliografía está en inglés.
- 2 de los 6 libros recomendados (Pervasive systems and ubiquitous computing y Ambient Intelligence) no están disponibles en la biblioteca de la universidad.
- En particular, el primer libro tiene un precio de 138€.

## User Journeys, Historias de Usuario y Milestones

En los siguientes documentos, se detallan los user journeys, historias de usuario y milestones que se han definido para el desarrollo de la aplicación Bibliofetch:

- [User Journeys](docs/user-journeys.md)
- [User Stories](docs/user-stories.md)
- [Milestones](docs/milestones.md)

## Configuración de Git y Github

- [Configuración de Git](docs/git_config.png)
- [Clave SSH de Github](docs/ssh_key.png)
