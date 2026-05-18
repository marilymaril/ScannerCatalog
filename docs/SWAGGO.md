# Swaggo Documentation Guide

## What is Swaggo?

Swaggo is a Go tool that automatically generates Swagger/OpenAPI documentation from Go code comments and type annotations. It parses special comment tags in your Go code and generates interactive API documentation.

## Installation

Install the Swag CLI:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Verify installation:

```bash
swag --version
```

## Generating Documentation

### Basic Usage

From the project root, point Swag to your main file:

```bash
swag init -g ./cmd/scannerApp/main.go
```

This generates three files in `docs/`:
- `docs.go` — Go package with embedded documentation
- `swagger.json` — OpenAPI specification in JSON
- `swagger.yaml` — OpenAPI specification in YAML

### Regenerate After Changes

Whenever you update Swagger comments or handlers, regenerate docs:

```bash
swag init -g ./cmd/scannerApp/main.go
```

## Package-Level Annotations

Add these comments at the top of your main file (before `package main`):

```go
// @title Scanner Catalog API
// @version 1.0
// @description This is a sample server for a Scanner Catalog API.
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
```

## Handler Function Annotations

Add comment tags directly above each HTTP handler function:

```go
// @Summary Get scanner item by ID
// @Description Get scanner item details by numeric ID.
// @Tags scanners
// @Accept json
// @Produce json
// @Param id path int true "Scanner ID"
// @Success 200 {object} Scanner
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /scanners/{id} [get]
func getItemByID(w http.ResponseWriter, r *http.Request) {
    // handler implementation
}
```

## Common Annotation Tags

| Tag | Purpose | Example |
|-----|---------|---------|
| `@Summary` | Brief description | `@Summary Get scanner by ID` |
| `@Description` | Detailed explanation | `@Description Retrieve scanner details...` |
| `@Tags` | Group related endpoints | `@Tags scanners` |
| `@Accept` | Input content type | `@Accept json` |
| `@Produce` | Output content type | `@Produce json` |
| `@Param` | Define parameters | `@Param id path int true "Scanner ID"` |
| `@Success` | Success response | `@Success 200 {object} Scanner` |
| `@Failure` | Error response | `@Failure 404 {object} ErrorResponse` |
| `@Router` | Endpoint route | `@Router /scanners/{id} [get]` |

## Parameter Syntax

Use `@Param` to define request parameters:

```go
// @Param id path int true "Scanner ID"
// @Param name query string false "Filter by name"
// @Param data body Scanner true "Scanner object"
```

Format: `@Param name location type required description`

Locations:
- `path` — URL path parameter (`:id`)
- `query` — Query string (`?name=value`)
- `body` — Request body
- `header` — HTTP header
- `formData` — Form data

## Response Syntax

Use `@Success` and `@Failure` to document responses:

```go
// @Success 200 {object} Scanner
// @Success 200 {array} Scanner
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
```

Format: `@Success code {type} Model`

Types:
- `{object}` — Single object
- `{array}` — Array of objects
- `{string}` — String response

## Type Definitions

Annotate your Go structs so Swaggo knows their fields:

```go
type Scanner struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Type string `json:"type"`
}

type ErrorResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}
```

The `json` tags are used to map Go fields to JSON keys.

## Viewing Documentation

### Option 1: Host with a Framework

For Gin:

```go
import "github.com/swaggo/gin-swagger"

r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

Visit: `http://localhost:8080/swagger/index.html`

### Option 2: Online Viewer

Use Swagger Editor online: https://editor.swagger.io/

Paste the contents of `docs/swagger.json` or upload the file.

### Option 3: View Raw Files

- `docs/swagger.json` — Read directly in a text editor or browser
- `docs/swagger.yaml` — Read directly in a text editor

## Common Issues

### Issue: "no Go files in ./"

**Cause**: Running `swag init` from a directory without Go files.

**Fix**: Specify the main file:
```bash
swag init -g ./cmd/scannerApp/main.go
```

### Issue: "cannot find type Package"

**Cause**: Swaggo can't find the Go package.

**Fix**: Ensure the main package file has the correct imports and package declaration.

### Issue: Docs not updating

**Cause**: Didn't regenerate docs after editing.

**Fix**: Re-run `swag init`:
```bash
swag init -g ./cmd/scannerApp/main.go
```

## Useful Commands

```bash
# Generate docs with verbose output
swag init -g ./cmd/scannerApp/main.go -v

# Generate docs in a specific directory
swag init -g ./cmd/scannerApp/main.go -o ./docs

# Format code in generated files
swag fmt
```

## Resources

- [Swaggo GitHub](https://github.com/swaggo/swag)
- [OpenAPI 3.0 Spec](https://swagger.io/specification/)
- [Swagger UI](https://swagger.io/tools/swagger-ui/)

## Next Steps

1. Uncomment the HTTP handler registration in `cmd/scannerApp/main.go`
2. Regenerate docs: `swag init -g ./cmd/scannerApp/main.go`
3. Start your server
4. View docs at the Swagger UI endpoint (if using a framework middleware)
