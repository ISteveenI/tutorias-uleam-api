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