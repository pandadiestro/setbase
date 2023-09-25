# SetBase
*the boolean database model*

`(ES)`

SetBase es una base de datos más conceptual que production-ready hecha en go basada en la idea de agrupar datos en sets, las queries están basadas en operaciones booleanas que se traducen como filtros para un grupo determinado de registros. Esto significa que, por ejemplo, `p & !q` significa "todos los registros que pertenezcan a `p` y no a `q`".

A modo de data de prueba estoy usando a la colección de waifus [este repositorio!](https://github.com/cat-milk/Anime-Girls-Holding-Programming-Books).

`(EN)`

SetBase is conceptual (not really production-ready) database made in Go based on the idea of grouping data into sets, the queries are based on Boolean operations that are translated as filters for a specific group of records. This means that, `p & !q`, e.g., means “all records that belong to p and not q”.

As a test data sample I'm using the waifus from [this amazing repo!](https://github.com/cat-milk/Anime-Girls-Holding-Programming-Books).

## Dependencies

`(ES)`

Cualquier toolit de Go, todas las librerías (paquetes) usadas son estándar.

`(EN)`

Some go toolkit, all used libraries (packages) are Go standard libraries.

## Preparing

`(ES)`

Se necesita especificar como variable de entorno el directorio donde se almacena todos los registros y su data, de aquí se leerá recursivamente entre directorios, esta variable es tal que:

`(EN)`

Before running the server you have to set up a local environment variable with the path of the data directory such that:

```sh
export setbase_data="./example_data/"
```

## Running

`(ES)`

Para activar el servidor local que corra en el puerto `:9090` del localhost

`(EN)`

To run the local server at the port `:9090` of localhost

```sh
go run setbase/src/main
```

## Querying

__ `(ES)`

Las queries se hacen, en este ejemplo, como peticiones web en:

`POST :9090/api -> []Register`

Donde el el cuerpo de la petición debe ser un objeto `JSON` tal que:

```go
type Query struct {
    Sets []string   `json:"sets"`
    Expr string     `json:"expr"`
}
```

Donde `sets` es una lista de las expresiones (tags, sets, etc.) que se van a usar en la expresión `expr`. Por ejemplo:

```JSON
{
    "sets": [
        "p",
        "q"
    ],
    "expr": "p && !q"
}
```

La respuesta de esta petición es una lista de registros donde cada registro tiene la estructura:

```go
type Register struct {
    Path string     `json:"path"`
    Sets []string   `json:"sets"`
}
```

__ `(EN)`

You can query the database, in this example server, at:

`POST :9090/api -> []Register`

Where the body of the request must be a `JSON` object such that:

```go
type Query struct {
    Sets []string   `json:"sets"`
    Expr string     `json:"expr"`
}
```

Where `sets` is a list of expressions (such as tags, sets, etc.) which will gonna be used at the `expr` expression, such as:

```JSON
{
    "sets": [
        "p",
        "q"
    ],
    "expr": "p && !q"
}
```

The response will also be a list of registers where each register has the following structure:

```go
type Register struct {
    Path string     `json:"path"`
    Sets []string   `json:"sets"`
}
```

## Media

__ `(ES)`

Todas las imágenes de este servidor de ejemplo son servidas estáticamente desde otro endpoint, en este caso es:

`
GET :9090/media?path=${path} -> bin
`

Donde cada registro de respuesta es enviado como data binaria cruda con el header de `Content-Type` como `image/png`

__ `(EN)`

The media of this example server (which statically serves all images from the example data directory) is retrieved from another endpoint, in this case:

`
GET :9090/media?path=${path} -> bin
`

Where Register is send with the `Content-Type` header set as `image/png`





