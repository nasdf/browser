//go:build !darwin && !windows

package browser

import (
	"context"
	"os/exec"
)

func Open(ctx context.Context, url string) error {
	_, err := exec.CommandContext(ctx, "xdg-open", url).CombinedOutput()
	return err
}
