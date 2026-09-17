//go:build unix

package docker

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/docker/docker/client"
)

// resizeExecTTY sets the guest PTY size to the host terminal and watches SIGWINCH.
// Without this, interactive TUIs often render but ignore keyboard input (size 0×0),
// and mouse-wheel scroll never attaches to the app.
func resizeExecTTY(ctx context.Context, cli client.ContainerAPIClient, execID string, inFd int) {
	doResize := func() {
		resizeExecOnce(ctx, cli, execID, inFd)
	}
	// Retry like docker CLI — resize can race with exec start.
	for i := 0; i < 10; i++ {
		doResize()
		time.Sleep(20 * time.Millisecond)
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	defer signal.Stop(ch)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ch:
			doResize()
		}
	}
}
