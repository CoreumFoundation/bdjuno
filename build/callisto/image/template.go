package image

import (
	"bytes"
	_ "embed"
	"text/template"
)

var (
	//go:embed Dockerfile.tmpl
	tmpl       string
	dockerfile = template.Must(template.New("dockerfile").Parse(tmpl))
)

// Data is the structure containing fields required by the template.
type Data struct {
	// From is the tag of the base image
	From string

	// Binary is the path to the binary file
	BinaryPath string

	// Binary is the name of faucet binary file to copy from build context
	BinaryName string

	// DBSchemaPath is the path to the database schema
	DBSchemaPath string
}

// Execute executes dockerfile template and returns complete dockerfile.
func Execute(data Data) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := dockerfile.Execute(buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
