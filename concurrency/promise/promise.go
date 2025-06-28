package promise

type Promise[T any] struct {
	Result <-chan T
	Error  <-chan error
}

func NewPromise[T any](operation func() (T, error)) *Promise[T] {
	r := make(chan T)
	e := make(chan error)
	p := &Promise[T]{
		Result: r,
		Error:  e,
	}

	go func() {
		defer close(r)
		defer close(e)

		res, err := operation()
		if err == nil {
			r <- res
		} else {
			e <- err
		}
	}()

	return p
}
