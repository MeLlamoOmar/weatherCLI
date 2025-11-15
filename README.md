**Nombre**: `omar/sun`

**Descripción**: Aplicación del clima escrita en Go. Consulta una "Weather API" remota para obtener datos meteorológicos (temperatura, humedad, descripción, etc.) y muestra la información por consola. La entrada está en `main.go` y la lógica auxiliar en `internals` (`config.go`, `types.go`). Usa `github.com/joho/godotenv` para cargar variables de entorno, `github.com/fatih/color` para salida coloreada y requiere una clave de API para la Weather API.

**Requisitos**:
- **Go**: versión `1.24.4` (según `go.mod`).
- **Entorno**: macOS / Linux / Windows con Go instalado.

**Instalación**:
- Clona el repositorio:

```bash
git clone <url-del-repo> && cd sun
```

- Descargar dependencias (opcional, `go` lo hará automáticamente al compilar):

```bash
go mod tidy
```

**Variables de entorno**:
- El proyecto usa `godotenv`. Crea un archivo `.env` en la raíz si necesitas configurar variables. Ejemplo mínimo:

Ajusta según lo que espere `internals/config.go`.

Se recomienda añadir la clave de la Weather API en el `.env` con la variable `WEATHER_API_KEY`. Ejemplo:

```env
# .env (ejemplo con API key)
WEATHER_API_KEY=tu_api_key_aqui
```

La aplicación usa la clave `WEATHER_API_KEY` para autenticar las peticiones a la Weather API. Revisa `internals/config.go` para ver nombres exactos y otros parámetros.

**Construir y ejecutar**:

- Ejecutar directamente con `go run`:

```bash
go run .
```

- Compilar binario:

```bash
go build -o bin/sun .
./bin/sun
```

**Estructura del proyecto**:
- `main.go`: punto de entrada de la aplicación.
- `internals/config.go`: carga/gestiona la configuración (variables de entorno, flags, etc.).
- `internals/types.go`: tipos y estructuras compartidas.
- `go.mod` / `go.sum`: dependencias y versión de Go.

**Dependencias notables**:
- `github.com/joho/godotenv`: carga variables de entorno desde `.env`.
- `github.com/fatih/color`: colorear la salida en terminal.

**Desarrollo**:
- Formateo del código:

```bash
gofmt -w .
```

- Ejecutar linters / herramientas (si las agrega más adelante).

**Pruebas**:
- Si hay tests, ejecuta:

```bash
go test ./...
```

**Notas**:
- Revisa `internals/config.go` para ver qué variables de entorno son necesarias y ajusta el `.env` en consecuencia.
- Si quieres que escriba una sección más detallada (API, flags, ejemplos de configuración) puedo inspeccionar `main.go` y `internals/` y extender el README.

**Próximos pasos sugeridos**:
- Añadir ejemplos concretos de `.env` según `internals/config.go`.
- Documentar los flags de línea de comandos si los hay.

---
Generado automáticamente por el asistente.