package desktop

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dsnet/try"
	"github.com/gosimple/slug"
	fm "github.com/liblaf/utils.go/pkg/std/fmt"
	o "github.com/liblaf/utils.go/pkg/std/os"
	"github.com/pterm/pterm"
)

type entry struct {
	Type string
	// Version              string
	Name string
	// GenericName          string
	// NoDisplay            bool
	// Comment              string
	Icon string
	// Hidden               string
	// OnlyShowIn           string
	// NotShowIn            string
	// DBusActivatable      bool
	// TryExec              string
	Exec string
	// Path                 string
	Terminal bool
	// Actions              string
	// MimeType             string
	// Categories           string
	// Implements           string
	// Keywords             string
	// StartupNotify        bool
	// StartupWMClass       string
	// URL                  string
	// PrefersNonDefaultGPU bool
	// SingleMainWindow     bool
}

func New(Type string, Name string) *entry {
	return &entry{
		Type: Type,
		Name: Name,
	}
}

func (e *entry) String() string {
	s := "[Desktop Entry]\n"
	s += KeyValueString("Type", e.Type)
	s += KeyValueString("Name", e.Name)
	s += KeyValueString("Icon", e.Icon)
	s += KeyValueString("Exec", e.Exec)
	s += KeyValueBool("Terminal", e.Terminal)
	return s
}

func KeyValueString(k string, v string) string {
	if len(v) == 0 {
		return ""
	}
	return fmt.Sprintf("%s=%s\n", k, v)
}

func KeyValueBool(k string, v bool) string {
	return fmt.Sprintf("%s=%t\n", k, v)
}

func (e *entry) Install() (err error) {
	defer try.Handle(&err)
	path := filepath.Join(UserApplications(), slug.Make(e.Name)+".desktop")
	try.E(os.WriteFile(path, []byte(e.String()), o.FilePerm()))
	pterm.Success.Println(fm.On("DESKTOP", path))
	return nil
}
