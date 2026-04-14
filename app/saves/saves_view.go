package saves

import (
	"after_the_end/app/components/backroundimage"
	"after_the_end/app/dialog/errorreport"
	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/backbone/styled"
	"after_the_end/db/model"

	"github.com/mappu/miqt/qt"
)

type View struct {
	*backbone.StatefullView[*Model]
	scrollArea *qt.QScrollArea
	listWidth  int
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
		Src:          ":/images/background.jpg",
		OverlayColor: "rgba(0, 0, 0, 0.6)",
	})

	widget.SetObjectName("saves")

	column := qt.NewQVBoxLayout2()
	column.SetObjectName("saves_column")

	column.AddStretch()
	column.AddWidget3(v.renderContainer(), 0, qt.AlignCenter)
	column.AddStretch()

	widget.Content.SetLayout(column.QLayout)
	return widget.QWidget
}

func (v *View) renderContainer() *qt.QWidget {
	screen := qt.QGuiApplication_PrimaryScreen().Geometry()
	width := min(int(float32(screen.Width())*0.6), 1000)
	height := min(int(float32(screen.Height())*0.6), 1000)

	widget := qt.NewQWidget2()
	widget.SetObjectName("saves_container")

	layout := qt.NewQVBoxLayout(widget)
	layout.SetObjectName("saves_container")
	layout.AddStretch()
	layout.AddWidget(v.renderTitle())

	v.scrollArea = qt.NewQScrollArea2()
	v.scrollArea.SetObjectName("saves_scroll")
	v.scrollArea.SetFixedSize2(width, height)
	v.scrollArea.SetStyleSheet(styled.S("#saves_scroll", styled.Transparent+"padding: 0"))
	v.scrollArea.VerticalScrollBar().SetStyleSheet(styled.CardScrollBar)

	v.listWidth = v.scrollArea.Width() - 32
	v.scrollArea.SetWidget(v.renderList())

	layout.AddWidget(v.scrollArea.QWidget)
	layout.AddWidget(v.renderBackButton())
	layout.AddStretch()
	return widget
}

func (v *View) renderTitle() *qt.QWidget {
	title := qt.NewQLabel3("Saves")
	title.SetStyleSheet(styled.Title2)
	title.SetGraphicsEffect(styled.TitleShadow())
	title.SetContentsMargins(0, 0, 0, 10)
	return title.QWidget
}

func (v *View) renderList() *qt.QWidget {
	v.list = NewListView(&ListViewOptions{
		GameSaves: v.Model.List,
		OnDelete:  v.deleteSave,
	})

	widget := v.Mount(v.list)
	widget.SetFixedWidth(v.listWidth)
	return widget
}

func (v *View) renderBackButton() *qt.QWidget {
	button := qt.NewQPushButton4(qt.NewQIcon4(":/icons/back-main.svg"), "Back")
	button.SetStyleSheet(styled.Button)
	button.SetFixedWidth(v.listWidth)

	button.OnReleased(func() {
		router.Push(router.RouteStart)
	})

	return button.QWidget
}

func (v *View) deleteSave(gameSave *model.GameSave) {
	if err := v.Model.Delete(gameSave); err != nil {
		errorreport.Show(v.Root, err)
		return
	}

	v.ViewUpdate()
}

func (v *View) ViewUpdate() {
	v.Unmount(v.list)
	v.scrollArea.SetWidget(v.renderList())
}
