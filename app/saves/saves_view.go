package saves

import (
	"after_the_end/app/components/backroundimage"
	"after_the_end/app/components/maincolumn"
	"after_the_end/app/dialog/reporterr"
	"after_the_end/app/resources"
	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/backbone/styled"
	"after_the_end/db/model"
	"after_the_end/helper/qtgeometry"

	"github.com/mappu/miqt/qt"
)

const (
	scrollSafeZone = 32
)

type View struct {
	*backbone.StatefullView[*Model]
	mainColumn *maincolumn.Widget
	scrollArea *qt.QScrollArea
	list       *ListView
}

func NewView() *View {
	return &View{
		StatefullView: backbone.NewStatefullView(NewModel()),
	}
}

func (v *View) ViewBeforeOpen(_ router.Params) error {
	return v.Model.Load()
}

func (v *View) ViewInit() *qt.QWidget {
	widget := backroundimage.New(&backroundimage.Options{
		Src:          resources.Image("background.jpg"),
		OverlayColor: backroundimage.OverlayDark,
	})

	widget.SetObjectName("saves")

	v.mainColumn = maincolumn.New(widget.Content)
	v.mainColumn.SetObjectName("saves_column")

	v.mainColumn.AddStretchWithStretch(1)
	v.mainColumn.AddWidget(v.renderTitle())
	v.mainColumn.AddWidget2(v.renderScrollArea(), 5)
	v.mainColumn.AddLayout(v.renderBackButton())
	v.mainColumn.AddStretchWithStretch(1)

	return widget.QWidget
}

func (v *View) renderTitle() *qt.QWidget {
	title := qt.NewQLabel3("Saves")
	title.SetStyleSheet(styled.Title2)
	title.SetGraphicsEffect(styled.TitleShadow())
	title.SetContentsMargins(0, 0, 0, 10)
	title.SetSizePolicy2(qt.QSizePolicy__Expanding, qt.QSizePolicy__Fixed)
	return title.QWidget
}

func (v *View) renderScrollArea() *qt.QWidget {
	v.scrollArea = qt.NewQScrollArea2()
	v.scrollArea.SetObjectName("saves_scroll")
	v.scrollArea.SetMaximumHeight(1000)
	v.scrollArea.SetStyleSheet(styled.S("#saves_scroll", styled.Transparent+"padding: 0"))
	v.scrollArea.VerticalScrollBar().SetStyleSheet(styled.CardScrollBar)
	v.scrollArea.SetWidget(v.renderList())
	return v.scrollArea.QWidget
}

func (v *View) renderList() *qt.QWidget {
	v.list = NewListView(&ListViewOptions{
		GameSaves: v.Model.List,
		OnDelete:  v.deleteSave,
	})

	widget := v.Mount(v.list)
	widget.SetContentsMargins(0, 0, scrollSafeZone, 0)

	qtgeometry.Read(v.mainColumn.Container, func(geometry *qt.QRect) {
		widget.SetFixedWidth(geometry.Width() - scrollSafeZone)
	})

	return widget
}

func (v *View) renderBackButton() *qt.QLayout {
	button := qt.NewQPushButton3("Back")
	button.SetStyleSheet(styled.Button)
	button.SetSizePolicy2(qt.QSizePolicy__Expanding, qt.QSizePolicy__Fixed)
	button.OnClicked(v.onBack)

	layout := qt.NewQHBoxLayout2()
	layout.SetContentsMargins(1, 0, scrollSafeZone+8, 0)
	layout.AddWidget(button.QWidget)

	return layout.QLayout
}

func (v *View) deleteSave(gameSave *model.GameSave) {
	if err := v.Model.Delete(gameSave); err != nil {
		reporterr.Show(v.ViewRoot(), err)
		return
	}

	if len(v.Model.List) == 0 {
		v.onBack()
		return
	}

	v.ViewUpdate()
}

func (v *View) ViewUpdate() {
	v.Unmount(v.list)
	v.scrollArea.SetWidget(v.renderList())
}

func (v *View) ViewDestroy() {
	v.StatefullView.ViewDestroy()
	v.mainColumn = nil
	v.scrollArea = nil
	v.list = nil
}

func (v *View) onBack() {
	router.Push(router.RouteStart)
}
