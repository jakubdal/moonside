package mt

import (
	"fmt"
	"strings"

	"github.com/go-test/deep"
)

func Equals(expected, got interface{}) error {
	diff := deep.Equal(expected, got)
	if diff != nil {
		return fmt.Errorf("expected != got;\n\t%s", strings.Join(diff, "\n\t"))
	}
	return nil
}
