package saves

import (
	"after_the_end/app/components/backroundimage"
	"after_the_end/app/components/maincolumn"
	"after_the_end/app/dialog/errorreport"
	"after_the_end/app/resources"
	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/backbone/styled"
	"after_the_end/db/model"
	"after_the_end/helper/qtgeometry"

	"github.com/mappu/miqt/qt"
)

type View struct {
	*backbone.StatefullView[*Model]
	mainColumn *maincolumn.Widget
	scrollArea *qt.QScrollArea
	list       *ListView
	backButton *qt.QPushButton
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

	v.mainColumn.AddStretch()
	v.mainColumn.AddWidget(v.renderTitle())
	v.mainColumn.AddWidget(v.renderScrollArea())
	v.mainColumn.AddWidget(v.renderBackButton())
	v.mainColumn.AddStretch()

	qtgeometry.Read(v.mainColumn.Container, v.resizeScrollArea)
	return widget.QWidget
}

func (v *View) renderTitle() *qt.QWidget {
	title := qt.NewQLabel3("Saves")
	title.SetStyleSheet(styled.Title2)
	title.SetGraphicsEffect(styled.TitleShadow())
	title.SetContentsMargins(0, 0, 0, 10)
	return title.QWidget
}

func (v *View) renderScrollArea() *qt.QWidget {
	v.scrollArea = qt.NewQScrollArea2()
	v.scrollArea.SetObjectName("saves_scroll")
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

	return v.Mount(v.list)
}

func (v *View) renderBackButton() *qt.QWidget {
	v.backButton = qt.NewQPushButton3("Back")
	v.backButton.SetStyleSheet(styled.Button)

	v.backButton.OnClicked(func() {
		router.Push(router.RouteStart)
	})

	return v.backButton.QWidget
}

func (v *View) resizeScrollArea(geometry *qt.QRect) {
	height := min(int(float32(geometry.Height())*0.6), 1000)
	width := geometry.Width()

	v.scrollArea.SetFixedSize2(width, height)

	scrollableWidth := width - 32
	v.list.ViewRoot().SetFixedWidth(scrollableWidth)
	v.backButton.SetFixedWidth(scrollableWidth)
}

func (v *View) deleteSave(gameSave *model.GameSave) {
	if err := v.Model.Delete(gameSave); err != nil {
		errorreport.Show(v.ViewRoot(), err)
		return
	}

	v.ViewUpdate()
}

func (v *View) ViewUpdate() {
	v.Unmount(v.list)
	v.scrollArea.SetWidget(v.renderList())
}
