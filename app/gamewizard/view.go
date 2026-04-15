package gamewizard

import (
	"after_the_end/app/components/backroundimage"
	"after_the_end/app/components/maincolumn"
	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/db/model"
	"after_the_end/helper/qtgeometry"

	"github.com/mappu/miqt/qt"
)

type View struct {
	*backbone.StatefullView[*Model]
	state      backbone.View
	mainColumn *maincolumn.Widget
}

func NewView() *View {
	return &View{
		StatefullView: backbone.NewStatefullView(NewModel()),
	}
}

func (v *View) ViewBeforeOpen(params router.Params) error {
	return v.Model.Load(params["gameSave"].(*model.GameSave))
}

func (v *View) ViewInit() *qt.QWidget {
	widget := backroundimage.New(&backroundimage.Options{
		Src:          ":/images/background.jpg",
		OverlayColor: backroundimage.OverlayDark,
	})

	widget.SetObjectName("wizard")

	v.mainColumn = maincolumn.New(widget.Content)
	v.renderState(widget.Content)

	qtgeometry.Read(widget.Content, func(geometry *qt.QRect) {
		v.state.ViewRoot().SetGeometryWithGeometry(geometry)
	})

	return widget.QWidget
}

func (v *View) ViewUpdate() {
	v.Unmount(v.state)
	v.renderState(v.ViewRoot())
}

func (v *View) renderState(parent *qt.QWidget) {
	switch v.Model.GameSave.State {
	case model.GameSaveNew:
		v.state = NewNameCharacterView(v.Model)
	}

	if v.state == nil {
		return
	}

	widget := v.Mount(v.state)
	widget.SetParent(parent)
	widget.SetGeometryWithGeometry(parent.Geometry())
}
