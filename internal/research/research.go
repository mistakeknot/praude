package research

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func Create(dir, id string, now time.Time) (string, error) {
	name := fmt.Sprintf("%s-%s.md", id, now.UTC().Format("20060102-150405"))
	path := filepath.Join(dir, name)
	body := fmt.Sprintf("# Research for %s\n\n- Competitive analysis:\n- Market summary:\n", id)
	return path, os.WriteFile(path, []byte(body), 0o644)
}
