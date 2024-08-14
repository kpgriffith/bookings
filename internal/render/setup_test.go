package render

import (
	"os"
	"testing"

	"github.com/alexedwards/scs/v2"
	"github.com/kpgriffith/bookings/internal/config"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}
