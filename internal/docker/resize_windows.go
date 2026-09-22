//go:build windows

package docker

import (
	"context"
	"time"

	"github.com/docker/docker/client"
)

// resizeExecTTY retries an initial resize; Windows has no SIGWINCH.
func resizeExecTTY(ctx context.Context, cli client.ContainerAPIClient, execID string, inFd int) {
	_ = ctx
	for i := 0; i < 10; i++ {
		resizeExecOnce(ctx, cli, execID, inFd)
		time.Sleep(20 * time.Millisecond)
	}
}
