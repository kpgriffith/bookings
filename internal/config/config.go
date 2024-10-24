package config

// avoid importing anything from the application.
// this should only import from the standard library to
// avoid cyclical dependency issues.
import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/kpgriffith/bookings/internal/models"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	MailChan      chan models.MailData
}
