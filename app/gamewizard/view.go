package gamewizard

import (
	"slices"

	"after_the_end/app/components/backroundimage"
	"after_the_end/app/components/maincolumn"
	"after_the_end/app/dialog/reporterr"
	"after_the_end/app/resources"
	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/db/model"

	qt "github.com/mappu/miqt/qt6"
)

type View struct {
	*backbone.StatefullView[*GameSaveModel]
	state      backbone.View
	mainColumn *maincolumn.Widget
	returnTo   router.Name
}

func NewView() *View {
	return &View{
		StatefullView: backbone.NewStatefullView(NewGameSaveModel()),
	}
}

func (v *View) ViewBeforeOpen(params router.Params) error {
	if returnTo, ok := params["returnTo"].(router.Name); ok {
		v.returnTo = returnTo
	} else {
		v.returnTo = router.RouteStart
	}

	return v.Model.Load(params["gameSave"].(*model.GameSave))
}

func (v *View) ViewInit() *qt.QWidget {
	widget := backroundimage.New(&backroundimage.Options{
		Src:          resources.Image("background.jpg"),
		OverlayColor: backroundimage.OverlayDark,
	})

	v.mainColumn = maincolumn.New(widget.Content)
	v.renderState()

	return widget.QWidget
}

func (v *View) ViewUpdate() {
	v.Unmount(v.state)
	v.renderState()
}

func (v *View) renderState() {
	switch v.Model.GameSave.State {
	case model.GameSaveNew:
		v.state = NewNameCharacterView(&NameCharacterViewOptions{
			Model:  v.Model,
			OnBack: v.onBack,
			OnNext: v.onNext,
		})
	case model.GameSaveGeneratingMap:
		v.state = NewGeneratingMapView(&GeneratingMapViewOptions{
			Model:  v.Model,
			OnNext: v.onNext,
		})
	}

	if v.state == nil {
		return
	}

	v.mainColumn.AddWidget(v.Mount(v.state))
}

func (v *View) onBack() {
	router.Push(v.returnTo)
}

func (v *View) onNext() {
	if err := v.nextState(); err != nil {
		reporterr.Show(v.ViewRoot(), err)
		return
	}

	if v.Model.GameSave.State == model.GameSaveReady {
		v.onComplete()
		return
	}

	v.ViewUpdate()
}

func (v *View) nextState() error {
	index := slices.Index(model.GameSaveStateVariants, v.Model.GameSave.State)
	v.Model.GameSave.State = model.GameSaveStateVariants[index+1]
	return v.Model.Save("state")
}

func (v *View) onComplete() {
	router.Push(router.RouteGame, router.Params{
		"gameSave": v.Model.GameSave,
	})
}

func (v *View) ViewDestroy() {
	v.StatefullView.ViewDestroy()
	v.mainColumn = nil
	v.state = nil
}
