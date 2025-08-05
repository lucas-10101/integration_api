package static

import (
	"embed"
)

//go:embed swagger
var SwaggerFS embed.FS

//go:embed translations
var TranslationsFS embed.FS
