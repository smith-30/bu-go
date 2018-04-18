package copier

import "github.com/otiai10/copy"

var Copy = func(src, dst string) error {
	err := copy.Copy(src, dst)
	return err
}
