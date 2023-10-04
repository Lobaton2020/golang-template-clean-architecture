TODO
- Dependency injection OK
- Load ENVS PENDING
- Create structure using modules PENDING
- Apply the use of interfaces for methods OK
- Libreria de validaciones de data OK
- Configurar un ORM (Opcional)
- Ver como embeber Swagger para cada endpoint OK
- Conversion - Validacion de data, params, queryParams, body, al handler OK
- Validacion de autenticacion, Autorizacion TOKEN (Sessiones - Gorilla libreria) NO
- Crear 2 modulos con get y paginacion OK
- Conexion con base de datos prosgresql, aplicar el uso de contexto OK
```
    create table products(
        id BIGSERIAL not null,
        name varchar(100) not null,
        price float not null,
        brand varchar(50) default 'generic',
        created_at timestamp default now(),
        primary key (id)
    );

    create table users(
        id BIGSERIAL not null,
        name varchar(100) not null,
        email varchar(100) not null,
        age int null,
        created_at timestamp default now() ,
        primary key (id)
    );
```
- Crear despliegue usndo dockerFile
- Configuracion Swagger PENDING

