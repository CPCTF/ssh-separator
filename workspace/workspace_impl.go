package workspace

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/errdefs"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/mazrean/separated-webshell/domain"
)

var (
	imageRef  string = os.Getenv("IMAGE_URL")
	imageUser string = os.Getenv("IMAGE_USER")
	opts             = types.ContainerAttachOptions{
		Stdin:  true,
		Stdout: true,
		Stderr: true,
		Stream: true,
	}
	containerMap = sync.Map{}
	stopTimeout  = 10 * time.Second
)

type containerInfo struct {
	id         string
	manageChan chan struct{}
}

func containerName(userName string) string {
	return fmt.Sprintf("user-%s", userName)
}

type Workspace struct {
	cli *client.Client
}

func NewWorkspace() (*Workspace, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create docker client: %w", err)
	}

	ctx := context.Background()

	reader, err := cli.ImagePull(ctx, imageRef, types.ImagePullOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to pull image: %w", err)
	}
	io.Copy(os.Stdout, reader)

	return &Workspace{
		cli: cli,
	}, nil
}

func (w *Workspace) Create(ctx context.Context, userName string) error {
	ctnName := containerName(userName)
	res, err := w.cli.ContainerCreate(ctx, &container.Config{
		Image:        imageRef,
		User:         imageUser,
		Tty:          true,
		OpenStdin:    true,
		AttachStderr: true,
		AttachStdin:  true,
		AttachStdout: true,
		StdinOnce:    true,
		Volumes:      make(map[string]struct{}),
	}, nil, nil, nil, ctnName)
	if errdefs.IsConflict(err) {
		ctnInfo, err := w.cli.ContainerInspect(ctx, ctnName)
		if err != nil {
			return fmt.Errorf("failed to inspect container: %w", err)
		}

		containerMap.Store(userName, &containerInfo{
			id:         ctnInfo.ID,
			manageChan: make(chan struct{}, 20),
		})

		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to create container: %w", err)
	}

	containerMap.Store(userName, &containerInfo{
		id:         res.ID,
		manageChan: make(chan struct{}, 20),
	})

	return nil
}

func (w *Workspace) Connect(ctx context.Context, userName string, isTty bool, winCh <-chan *domain.Window, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	iContainerInfo, ok := containerMap.Load(userName)
	if !ok {
		return errors.New("load container info error")
	}
	ctnInfo, ok := iContainerInfo.(*containerInfo)
	if !ok {
		return errors.New("parse container info error")
	}

	if len(ctnInfo.manageChan) >= 20 {
		return errors.New("too many shell")
	}

	err := w.cli.ContainerStart(ctx, ctnInfo.id, types.ContainerStartOptions{})
	if err != nil && !errdefs.IsConflict(err) {
		return fmt.Errorf("failed to start container: %w", err)
	}

	if isTty {
		go func() {
			for win := range winCh {
				err := w.cli.ContainerResize(ctx, ctnInfo.id, types.ResizeOptions{
					Height: win.Height,
					Width:  win.Width,
				})
				if err != nil {
					log.Println(err)
					break
				}
			}
		}()
	}

	ctnInfo.manageChan <- struct{}{}
	defer func(ctx context.Context, ctnInfo *containerInfo) {
		<-ctnInfo.manageChan
		if len(ctnInfo.manageChan) == 0 {
			err := w.cli.ContainerStop(ctx, ctnInfo.id, &stopTimeout)
			if err != nil {
				log.Fatalf("failed to stop container:%+v", err)
			}
		}
	}(ctx, ctnInfo)

	stream, err := w.cli.ContainerAttach(ctx, ctnInfo.id, opts)
	if err != nil {
		return fmt.Errorf("failed to attach container: %w", err)
	}
	defer stream.Close()

	outputErr := make(chan error)

	go func() {
		var err error
		if isTty {
			_, err = io.Copy(stdout, stream.Reader)
		} else {
			_, err = stdcopy.StdCopy(stdout, stderr, stream.Reader)
		}
		outputErr <- err
	}()

	go func() {
		defer stream.CloseWrite()
		io.Copy(stream.Conn, stdin)
	}()

	resultC, errC := w.cli.ContainerWait(ctx, ctnInfo.id, container.WaitConditionNotRunning)
	select {
	case err = <-errC:
		return fmt.Errorf("failed to wait container: %w", err)
	case <-resultC:
	}
	err = <-outputErr
	return nil
}

func (*Workspace) Remove(ctx context.Context, userName string) error {
	return nil
}