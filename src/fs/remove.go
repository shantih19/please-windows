package fs

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// ForceRemoveAll will attempt an os.RemoveAll() but then resort to a more forceful method:
// 1) chmod the files and folders to make sure we have permissions (we need +rwx on dirs even with rm -rf)
// 2) rm -rf path
func ForceRemoveAll(path string) error {
	if err := os.RemoveAll(path); err != nil {
		cmd := exec.Command("bash", "-c", fmt.Sprintf("chmod -R +rwx %s; rm -rf %s", path, path))

		buf := new(bytes.Buffer)

		cmd.Stdout = buf
		cmd.Stderr = buf

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("%w\n%s", err, buf.String())
		}
		return nil
	}
	return nil
}