# API Backend para Gestión de Tutorías Académicas ULEAM

Este proyecto consiste en una API REST desarrollada en Go para gestionar tutorías académicas entre estudiantes y docentes de la Universidad Laica Eloy Alfaro de Manabí.

## Problema

Actualmente la coordinación de tutorías académicas puede realizarse mediante mensajes, conversaciones presenciales o acuerdos verbales, generando problemas como confusión de horarios, falta de seguimiento de solicitudes, dificultad para organizar sesiones y pérdida de evidencias de las tutorías realizadas.

## Solución

La API permitirá gestionar el proceso completo de tutorías académicas mediante módulos independientes:

- Los docentes podrán registrar sus materias y horarios disponibles.
- Los estudiantes podrán solicitar tutorías indicando el tema a tratar y tipo de tutoría.
- El sistema permitirá gestionar sesiones realizadas, asistencia y evidencias.

## Módulos

### 1. Módulo Docentes

Gestiona la información de los docentes, materias asignadas y horarios disponibles para tutorías.

Entidades:

- Docente
- Materia
- HorarioDocente


### 2. Módulo Solicitudes de Tutoría

Gestiona las solicitudes realizadas por estudiantes.

Entidades:

- Estudiante
- TipoTutoria
- SolicitudTutoria


### 3. Módulo Sesiones de Tutoría

Gestiona las tutorías confirmadas, asistencia y evidencias generadas.

Entidades:

- SesionTutoria
- Asistencia
- Evidencia


## Stack tecnológico

- Go
- Chi Router
- GORM
- JWT
- Testify
- Docker
- SQLite para desarrollo
- PostgreSQL para producción
- GitHub Actions


## Integrantes y módulos

- Steveen Acosta: estructura inicial del proyecto y módulo Sesiones de Tutoría.
- Karen Holguín: módulo Docentes.
- Jorge Mero: módulo Solicitudes de Tutoría.


# Módulo Docentes

Este módulo permite administrar la información relacionada con los docentes, materias y horarios destinados para la gestión de tutorías académicas.

## Entidades implementadas

### Docente

Permite registrar la información del docente encargado de brindar tutorías.

Campos principales:

- id
- nombres
- apellidos
- correo
- teléfono
- departamento
- especialidad
- título académico


### Materia

Permite gestionar las asignaturas relacionadas con las tutorías.

Campos principales:

- id
- nombre
- código
- descripción


### HorarioDocente

Permite registrar los horarios disponibles de cada docente asociados a una materia.

Campos principales:

- id
- docente_id
- materia_id
- día_semana
- hora_inicio
- hora_fin
- modalidad
- aula


## Funcionalidades implementadas

- Registrar docentes.
- Consultar docentes.
- Actualizar información de docentes.
- Eliminar docentes.
- Registrar materias.
- Consultar materias.
- Registrar horarios docentes.
- Consultar horarios disponibles.


## Endpoints

### Docentes

POST

/api/v1/docentes


GET

/api/v1/docentes


GET por ID

/api/v1/docentes/{id}


PUT

/api/v1/docentes/{id}


DELETE

/api/v1/docentes/{id}



### Materias

POST

/api/v1/materias


GET

/api/v1/materias



### Horarios

POST

/api/v1/horarios


GET

/api/v1/horarios


## Regla de negocio

Un docente no puede registrar dos horarios de tutoría que se crucen en el mismo día y rango de horas.