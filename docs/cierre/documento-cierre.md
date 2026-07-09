\# Documento de cierre - Hito 3



\## Proyecto



El proyecto desarrollado es una API backend para la gestion de tutorias academicas de la ULEAM. El sistema permite administrar docentes, materias, horarios, estudiantes, solicitudes de tutoria, sesiones, asistencias y evidencias.



\## Que aprendimos



Durante el desarrollo aprendimos a estructurar una API REST en Go usando arquitectura por capas. La capa handler recibe las peticiones HTTP, la capa service concentra reglas de negocio y validaciones, y la capa repository gestiona la persistencia mediante GORM.



Tambien aprendimos a trabajar con Docker y PostgreSQL para levantar el proyecto completo en un entorno reproducible. El sistema se puede ejecutar con docker compose, conectando la API con la base de datos y cargando datos iniciales mediante seeders.



Otro aprendizaje importante fue el uso de GitHub Actions para ejecutar automaticamente build, vet y tests. Esto nos permitio validar los cambios antes de integrarlos a main mediante Pull Requests y proteccion de rama.



En el modulo de sesiones de tutoria se reforzaron pruebas en las tres capas: handler, service y repository. Se usaron mocks con testify, httptest para rutas HTTP y SQLite en memoria para pruebas de persistencia.



\## Que hariamos distinto



En una siguiente version organizariamos el proyecto usando Vertical Slice, separando cada modulo en su propia carpeta con sus modelos, handlers, services, repositories y tests. Esto reduciria conflictos entre integrantes y facilitaria la revision por Pull Request.



Tambien implementariamos desde el inicio un modulo completo de autenticacion JWT con usuarios, login, registro y roles. Ademas, documentariamos los endpoints desde las primeras semanas para no concentrar la documentacion al final.



\## Proximos pasos



Los proximos pasos del producto son completar la autenticacion JWT con roles, mejorar los permisos por usuario, agregar filtros por fecha, docente y estudiante, generar reportes de tutorias y preparar el sistema para un despliegue en un ambiente real.



Tambien se recomienda mantener la rama main protegida, continuar usando Pull Requests con checks obligatorios y aumentar la cobertura de pruebas en todos los modulos del proyecto.

