package adapter

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/toolbox"
	"path"
	"testing"
)

func TestGenerator_Generate(t *testing.T) {

	parentDir := toolbox.CallerDirectory(3)
	gen := New()

	code, err := gen.Generate(path.Join(parentDir, "test"), "MyInterface", func(receiver *toolbox.FunctionInfo) bool {
		return true
	}, func(metaType *TypeMeta, receiver *toolbox.FunctionInfo) {

	})
	assert.Nil(t, err)
	assert.NotNil(t, code)
}
