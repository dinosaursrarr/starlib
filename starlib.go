package starlib

import (
	"fmt"

	"github.com/qri-io/starlib/encoding/base64"
	"github.com/qri-io/starlib/html"
	"github.com/qri-io/starlib/http"
	"github.com/qri-io/starlib/re"
	"github.com/qri-io/starlib/time"
	"github.com/qri-io/starlib/xlsx"
	"github.com/qri-io/starlib/zipfile"
	"go.starlark.net/starlark"
)

// Loader presents the starlib library as a loader
func Loader(thread *starlark.Thread, module string) (dict starlark.StringDict, err error) {
	switch module {
	case time.ModuleName:
		return time.LoadModule()
	case http.ModuleName:
		return http.LoadModule()
	case xlsx.ModuleName:
		return xlsx.LoadModule()
	case html.ModuleName:
		return html.LoadModule()
	case zipfile.ModuleName:
		return zipfile.LoadModule()
	case re.ModuleName:
		return re.LoadModule()
	case base64.ModuleName:
		return base64.LoadModule()
	}

	return nil, fmt.Errorf("invalid module")
}
