package utils

import (
	"context"
	"errors"
	"io"
)

func Copy(ctx context.Context, src io.Reader, dst io.Writer) error {
	buffer := make([]byte, 32*1024)

	var err error
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		_, err = src.Read(buffer)
		if err != nil {
			break
		}

		_, err = dst.Write(buffer)
		if err != nil {
			break
		}
	}

	if errors.Is(err, io.EOF) {
		return nil
	}

	return err
}

func CopyFile(ctx context.Context, src, dst string) error {
	return nil
}

func CopyWithProgress(ctx context.Context, src io.Reader, dst io.Reader) (<-chan int, <-chan error) {
	return nil, nil
}

func CopyFileWithProgress(ctx context.Context, src, dst string) (<-chan int, <-chan error) {
	return nil, nil
}
