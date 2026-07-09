\# Diagrama de arquitectura - API de Tutorias ULEAM



\## Arquitectura general



```text

Cliente / Postman / Frontend

&#x20;       |

&#x20;       v

Router Chi

&#x20;       |

&#x20;       v

Handler

&#x20;       |

&#x20;       v

Service

&#x20;       |

&#x20;       v

Repository

&#x20;       |

&#x20;       v

GORM

&#x20;       |

&#x20;       v

PostgreSQL

Responsabilidad de las capas

Handler



Recibe las peticiones HTTP, lee parametros de URL, decodifica JSON, llama al service y responde con codigos HTTP y JSON.



Ejemplo en sesiones:



SesionHandler.Create

SesionHandler.GetAll

SesionHandler.GetByID

SesionHandler.Update

SesionHandler.Delete

SesionHandler.CreateAsistencia

SesionHandler.CreateEvidencia

Service



Contiene reglas de negocio y validaciones. No conoce detalles HTTP ni SQL.



Ejemplo en sesiones:



Validar solicitud\_id

Validar fecha\_sesion

Validar hora\_inicio y hora\_fin

Asignar estado por defecto

Validar asistencia

Validar evidencia

Repository



Encapsula el acceso a datos con GORM.



Ejemplo en sesiones:



Create

FindByID

FindAll

Update

Delete

CreateAsistencia

CreateEvidencia

Integracion entre modulos

SolicitudTutoria

&#x20;       |

&#x20;       | solicitud\_id

&#x20;       v

SesionTutoria

&#x20;       |

&#x20;       +---- Asistencia

&#x20;       |

&#x20;       +---- Evidencia



La sesion de tutoria se relaciona con una solicitud mediante solicitud\_id. A partir de una sesion se registran asistencias y evidencias. En GORM se aplica una relacion Has-Many desde SesionTutoria hacia Asistencia y Evidencia.



Flujo de una request

POST /api/v1/sesiones-tutoria/

&#x20;       |

&#x20;       v

RequireAuth

&#x20;       |

&#x20;       v

SesionHandler.Create

&#x20;       |

&#x20;       v

SesionService.Create

&#x20;       |

&#x20;       v

validarSesion

&#x20;       |

&#x20;       v

SesionRepository.Create

&#x20;       |

&#x20;       v

GORM + PostgreSQL



Guarda y cierra.



\### 3. Crear colección Postman básica



Para no depender de exportar manualmente, crea una colección JSON válida. Ejecuta:



```powershell

notepad docs\\postman\\tutorias-uleam-api.postman\_collection.json

{

&#x20; "info": {

&#x20;   "name": "Tutorias ULEAM API",

&#x20;   "description": "Coleccion Postman del proyecto API Backend para gestion de tutorias academicas ULEAM.",

&#x20;   "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"

&#x20; },

&#x20; "item": \[

&#x20;   {

&#x20;     "name": "Sesiones de Tutoria",

&#x20;     "item": \[

&#x20;       {

&#x20;         "name": "Listar sesiones",

&#x20;         "request": {

&#x20;           "method": "GET",

&#x20;           "url": "http://localhost:8080/api/v1/sesiones-tutoria/"

&#x20;         }

&#x20;       },

&#x20;       {

&#x20;         "name": "Obtener sesion por ID",

&#x20;         "request": {

&#x20;           "method": "GET",

&#x20;           "url": "http://localhost:8080/api/v1/sesiones-tutoria/1"

&#x20;         }

&#x20;       },

&#x20;       {

&#x20;         "name": "Crear sesion",

&#x20;         "request": {

&#x20;           "method": "POST",

&#x20;           "header": \[

&#x20;             {

&#x20;               "key": "Content-Type",

&#x20;               "value": "application/json"

&#x20;             },

&#x20;             {

&#x20;               "key": "Authorization",

&#x20;               "value": "Bearer token-prueba"

&#x20;             }

&#x20;           ],

&#x20;           "body": {

&#x20;             "mode": "raw",

&#x20;             "raw": "{\\n  \\"solicitud\_id\\": 1,\\n  \\"fecha\_sesion\\": \\"2026-07-01\\",\\n  \\"hora\_inicio\\": \\"09:00\\",\\n  \\"hora\_fin\\": \\"10:00\\",\\n  \\"observaciones\\": \\"Sesion creada desde Postman\\",\\n  \\"estado\\": \\"Programada\\"\\n}"

&#x20;           },

&#x20;           "url": "http://localhost:8080/api/v1/sesiones-tutoria/"

&#x20;         }

&#x20;       },

&#x20;       {

&#x20;         "name": "Registrar asistencia",

&#x20;         "request": {

&#x20;           "method": "POST",

&#x20;           "header": \[

&#x20;             {

&#x20;               "key": "Content-Type",

&#x20;               "value": "application/json"

&#x20;             },

&#x20;             {

&#x20;               "key": "Authorization",

&#x20;               "value": "Bearer token-prueba"

&#x20;             }

&#x20;           ],

&#x20;           "body": {

&#x20;             "mode": "raw",

&#x20;             "raw": "{\\n  \\"estudiante\_asistio\\": true,\\n  \\"docente\_asistio\\": true,\\n  \\"observacion\\": \\"Asistieron estudiante y docente\\"\\n}"

&#x20;           },

&#x20;           "url": "http://localhost:8080/api/v1/sesiones-tutoria/1/asistencias"

&#x20;         }

&#x20;       },

&#x20;       {

&#x20;         "name": "Registrar evidencia",

&#x20;         "request": {

&#x20;           "method": "POST",

&#x20;           "header": \[

&#x20;             {

&#x20;               "key": "Content-Type",

&#x20;               "value": "application/json"

&#x20;             },

&#x20;             {

&#x20;               "key": "Authorization",

&#x20;               "value": "Bearer token-prueba"

&#x20;             }

&#x20;           ],

&#x20;           "body": {

&#x20;             "mode": "raw",

&#x20;             "raw": "{\\n  \\"tipo\_archivo\\": \\"pdf\\",\\n  \\"archivo\_url\\": \\"https://example.com/evidencia.pdf\\",\\n  \\"descripcion\\": \\"Evidencia de la sesion\\"\\n}"

&#x20;           },

&#x20;           "url": "http://localhost:8080/api/v1/sesiones-tutoria/1/evidencias"

&#x20;         }

&#x20;       }

&#x20;     ]

&#x20;   },

&#x20;   {

&#x20;     "name": "Solicitudes de Tutoria",

&#x20;     "item": \[

&#x20;       {

&#x20;         "name": "Listar solicitudes",

&#x20;         "request": {

&#x20;           "method": "GET",

&#x20;           "url": "http://localhost:8080/api/v1/solicitudes-tutoria"

&#x20;         }

&#x20;       },

&#x20;       {

&#x20;         "name": "Obtener solicitud por ID",

&#x20;         "request": {

&#x20;           "method": "GET",

&#x20;           "url": "http://localhost:8080/api/v1/solicitudes-tutoria/1"

&#x20;         }

&#x20;       }

&#x20;     ]

&#x20;   },

&#x20;   {

&#x20;     "name": "Docentes",

&#x20;     "item": \[

&#x20;       {

&#x20;         "name": "Listar docentes",

&#x20;         "request": {

&#x20;           "method": "GET",

&#x20;           "url": "http://localhost:8080/api/v1/docentes"

&#x20;         }

&#x20;       },

&#x20;       {

&#x20;         "name": "Obtener docente por ID",

&#x20;         "request": {

&#x20;           "method": "GET",

&#x20;           "url": "http://localhost:8080/api/v1/docentes/1"

&#x20;         }

&#x20;       }

&#x20;     ]

&#x20;   }

&#x20; ]

}



