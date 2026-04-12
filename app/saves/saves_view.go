package saves

import (
	"log"

	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/backbone/styled"

	"github.com/mappu/miqt/qt"
)

type View struct {
	*backbone.StatefullView[*Model]
	layout *qt.QLayout
}

func NewView() *View {
	return &View{
		StatefullView: backbone.NewStatefullView(NewModel()),
	}
}

func (v *View) Layout() *qt.QLayout {
	return v.layout
}

func (v *View) ViewInit(parent *qt.QWidget) {
	v.Model.Load()

	widget := qt.NewQWidget2()
	widget.SetObjectName("saves")
	widget.SetStyleSheet("background: url(:/images/background.jpg)")

	column := qt.NewQVBoxLayout2()
	column.SetObjectName("saves_column")

	column.AddStretch()
	column.AddWidget3(v.renderContainer(), 0, qt.AlignCenter)
	column.AddStretch()

	widget.SetLayout(column.QLayout)

	cover := qt.NewQVBoxLayout(parent)
	cover.SetObjectName("start_window_cover")
	cover.SetContentsMargins(0, 0, 0, 0)
	cover.AddWidget(widget)
	v.layout = cover.QLayout
}

func (v *View) renderContainer() *qt.QWidget {
	screen := qt.QGuiApplication_PrimaryScreen().Geometry()
	width := min(int(float32(screen.Width())*0.6), 1000)
	height := min(int(float32(screen.Height())*0.6), 1000)

	log.Printf("screen width: %v, height: %v", width, height)

	widget := qt.NewQWidget2()
	widget.SetStyleSheet(styled.Reset)
	widget.SetObjectName("saves_container")

	layout := qt.NewQVBoxLayout(widget)
	layout.SetObjectName("saves_container")
	layout.AddStretch()
	layout.AddWidget(v.renderTitle())

	scrollArea := qt.NewQScrollArea2()
	scrollArea.SetObjectName("saves_scroll")
	scrollArea.SetFixedSize2(width, height)
	scrollArea.SetStyleSheet(styled.S("#saves_scroll", styled.Card))
	scrollArea.VerticalScrollBar().SetStyleSheet(styled.CardScrollBar)
	scrollArea.SetWidget(v.renderList(scrollArea))

	layout.AddWidget(scrollArea.QWidget)
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

func (v *View) renderList(scrollArea *qt.QScrollArea) *qt.QWidget {
	widget := qt.NewQWidget2()
	widget.SetStyleSheet(styled.Reset)
	widget.SetFixedWidth(scrollArea.Width() - scrollArea.VerticalScrollBar().Width())

	column := qt.NewQVBoxLayout(widget)
	for _, save := range v.Model.List {
		column.AddWidget(v.MountForLayout(NewSaveView(save)))
	}

	return widget
}

func (v *View) renderBackButton() *qt.QWidget {
	button := qt.NewQPushButton4(qt.NewQIcon4(":/icons/back-main.svg"), "Back")
	button.SetStyleSheet(styled.Button)

	button.OnReleased(func() {
		router.Push(router.RouteStart)
	})

	return button.QWidget
}
