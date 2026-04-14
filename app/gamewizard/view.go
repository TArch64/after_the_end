package gamewizard

import (
	"after_the_end/app/components/backroundimage"
	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/db/model"

	"github.com/mappu/miqt/qt"
)

type View struct {
	*backbone.StatefullView[*Model]
	state backbone.View
}

func NewView() *View {
	return &View{
		StatefullView: backbone.NewStatefullView(NewModel()),
	}
}

func (v *View) ViewBeforeOpen(params router.Params) error {
	v.Model.Load(params["gameSave"].(*model.GameSave))
	return nil
}

func (v *View) ViewInit() *qt.QWidget {
	widget := backroundimage.New(&backroundimage.Options{
		Src:          ":/images/background.jpg",
		OverlayColor: "rgba(0, 0, 0, 0.6)",
	})

	v.renderState(widget.QWidget)
	widget.OnResizeEvent(v.onResize)

	return widget.QWidget
}

func (v *View) ViewUpdate() {
	v.Unmount(v.state)
	v.renderState(v.ViewRoot())
}

func (v *View) renderState(parent *qt.QWidget) {
	switch v.Model.GameSave.State {
	case model.GameSaveCreateMainCharacter:
		v.state = NewCreateCharacterView(v.Model)
	}

	if v.state == nil {
		return
	}

	widget := v.Mount(v.state)
	widget.SetParent(parent)
	widget.SetGeometryWithGeometry(parent.Geometry())
}

func (v *View) onResize() {
	v.state.ViewRoot().SetGeometryWithGeometry(v.ViewRoot().Geometry())
}
