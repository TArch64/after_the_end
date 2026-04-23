package gamewizard

import (
	"time"

	"after_the_end/app/dialog/errorreport"
	"after_the_end/backbone"
	"after_the_end/backbone/styled"

	"github.com/mappu/miqt/qt"
	"github.com/mappu/miqt/qt/mainthread"
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
	widget.SetObjectName("generating_map")

	column := qt.NewQVBoxLayout(widget)
	column.SetObjectName("generating_map")

	column.AddStretch()
	column.AddWidget3(v.renderTitle(), 0, qt.AlignCenter)
	column.AddStretch()

	go v.generateMap()

	return widget
}

func (v *GeneratingMapView) renderTitle() *qt.QWidget {
	title := qt.NewQLabel3("Generating Random Map")
	title.SetObjectName("generating_map_title")
	title.SetStyleSheet(styled.S("#generating_map_title", styled.Title2))
	title.SetGraphicsEffect(styled.TitleShadow())
	return title.QWidget
}

func (v *GeneratingMapView) generateMap() {
	time.Sleep(1 * time.Second)
	err := v.model.GenerateMap()

	mainthread.Start(func() {
		if err != nil {
			errorreport.Show(v.ViewRoot(), err)
			return
		}

		v.onNext()
	})
}
