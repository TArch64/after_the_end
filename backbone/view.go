package backbone

import (
	"errors"
)

type View interface {
	ViewInit() error
	ViewUpdate()
	ViewDestroy() error
}

type BaseView struct {
	children []View
}

func NewBaseView() *BaseView {
	return &BaseView{}
}

func (b *BaseView) Mount(view View) error {
	if err := view.ViewInit(); err != nil {
		return err
	}

	b.children = append(b.children, view)
	return nil
}

func (b *BaseView) ViewDestroy() error {
	var errs []error
	for _, child := range b.children {
		if err := child.ViewDestroy(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
