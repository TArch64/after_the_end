package gamewizard

import (
	"time"

	"after_the_end/app/dialog/reporterr"
	"after_the_end/backbone"
	"after_the_end/backbone/styled"

	qt "github.com/mappu/miqt/qt6"
	"github.com/mappu/miqt/qt6/mainthread"
)

type GeneratingMapView struct {
	*backbone.StatelessView
	model  *GameSaveModel
	onNext func()
}

type GeneratingMapViewOptions struct {
	Model  *GameSaveModel
	OnNext func()
}

func NewGeneratingMapView(options *GeneratingMapViewOptions) *GeneratingMapView {
	return &GeneratingMapView{
		StatelessView: backbone.NewStatelessView(),
		model:         options.Model,
		onNext:        options.OnNext,
	}
}

func (v *GeneratingMapView) ViewInit() *qt.QWidget {
	widget := qt.NewQWidget2()
	column := qt.NewQVBoxLayout(widget)

	column.AddStretch()
	column.AddWidget3(v.renderTitle(), 0, qt.AlignCenter)
	column.AddStretch()

	go v.generateMap()

	return widget
}

func (v *GeneratingMapView) renderTitle() *qt.QWidget {
	title := qt.NewQLabel3("Generating Random Map")
	title.SetProperty("text-title", qt.NewQVariant4(2))
	title.SetGraphicsEffect(styled.TitleShadow())
	return title.QWidget
}

func (v *GeneratingMapView) generateMap() {
	time.Sleep(1 * time.Second)
	err := v.model.GenerateMap()

	mainthread.Start(func() {
		if err != nil {
			reporterr.Show(v.ViewRoot(), err)
			return
		}

		v.onNext()
	})
}
