package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "name") {
		return errors.New("name already taken")
	}
	return errors.New("incorrect details")
}
