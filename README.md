# API Backend para Gestión de Tutorías Académicas ULEAM

Este proyecto consiste en una API REST desarrollada en Go para gestionar turnos de tutorías académicas entre estudiantes y docentes de la Universidad Laica Eloy Alfaro de Manabí.

## Problema

La coordinación de tutorías académicas puede realizarse por mensajes, conversaciones presenciales o acuerdos verbales, lo que puede generar confusión sobre horarios, temas, confirmaciones y reprogramaciones.

## Solución

La API permitirá que los docentes registren su disponibilidad, que los estudiantes soliciten tutorías indicando materia, tema, urgencia y modalidad, y que el sistema gestione sesiones confirmadas evitando cruces de horario.

## Módulos

1. Disponibilidad Docente
2. Solicitudes de Tutoría
3. Sesiones de Tutoría

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

- Steveen Acosta: estructura inicial del proyecto.
- Karen Holguín: módulo Disponibilidad Docente.
- Jorge Mero: módulo Solicitudes de Tutoría.
- Steveen Acosta módulo Sesiones de Tutoría.


## Módulo Disponibilidad Docente

Este módulo permite registrar y administrar la disponibilidad horaria de los docentes para las tutorías académicas.

### Funcionalidades implementadas

* Registrar disponibilidad docente.

* Consultar todas las disponibilidades registradas.

* Consultar disponibilidad por identificador.

* Actualizar disponibilidad existente.

* Eliminar disponibilidad registrada.

* Validar datos obligatorios antes de registrar o actualizar información.

### Endpoints

#### Crear disponibilidad

POST /api/v1/disponibilidades

#### Obtener todas las disponibilidades

GET /api/v1/disponibilidades

#### Obtener disponibilidad por ID

GET /api/v1/disponibilidades/{id}

#### Actualizar disponibilidad

PUT /api/v1/disponibilidades/{id}

#### Eliminar disponibilidad

DELETE /api/v1/disponibilidades/{id}


## Módulo Solicitudes de Tutoría

Este módulo permite a los estudiantes registrar solicitudes de tutoria indicando la materia, el tema de consulta, el nivel de urgencia y la modalidad requerida.

### Funcionalidades implementadas

* Registrar solicitudes de tutoría.
* Consultar solicitudes registradas.
* Consultar una solicitud por identificador.
* Actualizar información de una solicitud existente.
* Eliminar solicitudes registradas.

### Endpoints

#### Crear solicitud

POST /api/v1/solicitudes-tutoria

#### Obtener todas las solicitudes

GET /api/v1/solicitudes-tutoria

#### Obtener solicitud por ID

GET /api/v1/solicitudes-tutoria/{id}

#### Actualizar solicitud

PUT /api/v1/solicitudes-tutoria/{id}

#### Eliminar solicitud

DELETE /api/v1/solicitudes-tutoria/{id}

