
import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

// ASSERT OMIT
func TestSomething(t *testing.T) {
	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	assert.Nil(t, object)

	// assert for not nil (good when you expect something)
	if assert.NotNil(t, object) {
		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "Something", object.Value)
	}
}

// END OMIT

// REQUIRE OMIT
func TestSomething(t *testing.T) {
	// assert equality
	require.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	require.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	require.Nil(t, object)

	// assert for not nil (good when you expect something)
	if require.NotNil(t, object) {
		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		require.Equal(t, "Something", object.Value)
	}
}

// END OMIT

// AST OMIT
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

func noErr(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

// END OMIT