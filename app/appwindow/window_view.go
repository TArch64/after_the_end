package appwindow

import (
	"os"

	"after_the_end/backbone"

	"github.com/mappu/miqt/qt"
)

type WindowView struct {
	*backbone.BaseView
}

func NewWindowView() *WindowView {
	return &WindowView{
		BaseView: backbone.NewBaseView(),
	}
}

func (v *WindowView) ViewInit() error {
	qt.NewQApplication(os.Args)

	if err := v.Mount(NewButtonView()); err != nil {
		return err
	}

	qt.QApplication_Exec()
	return nil
}

func (v *WindowView) ViewUpdate() {}
