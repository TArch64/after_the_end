package backbone

type Disposable interface {
	Dispose()
}

type DisposableController struct {
	handlers []Disposable
}

func NewDisposableController() *DisposableController {
	return &DisposableController{}
}

func (c *DisposableController) AutoDispose(disposable Disposable) {
	c.handlers = append(c.handlers, disposable)
}

func (c *DisposableController) dispose() {
	for _, handler := range c.handlers {
		handler.Dispose()
	}

	c.handlers = nil
}
