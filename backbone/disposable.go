package backbone

type Disposable interface {
	Dispose()
}

type DisposableView struct {
	handlers []Disposable
}

func NewDisposableView() *DisposableView {
	return &DisposableView{}
}

func (v *DisposableView) AutoDispose(disposable Disposable) {
	v.handlers = append(v.handlers, disposable)
}

func (v *DisposableView) dispose() {
	for _, handler := range v.handlers {
		handler.Dispose()
	}

	v.handlers = nil
}
