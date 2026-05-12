package game

import (
	"after_the_end/app/dialog/reporterr"
	"after_the_end/app/game/command"
	"after_the_end/app/game/command/cmd"
	"after_the_end/app/game/scene"
	"after_the_end/app/game/state"
	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/db/model"

	qt "github.com/mappu/miqt/qt6"
)

type View struct {
	*backbone.StatefullView[*state.Model]
}

func NewView() *View {
	return &View{
		StatefullView: backbone.NewStatefullView(state.NewModel()),
	}
}

func (v *View) ViewBeforeOpen(params router.Params) error {
	return v.Model.Load(params["gameSave"].(*model.GameSave))
}

func (v *View) ViewInit() *qt.QWidget {
	widget := qt.NewQWidget2()
	stack := qt.NewQStackedLayout(widget)
	// stack.AddWidget(v.Mount(overlay.NewView()))
	stack.AddWidget(v.Mount(scene.NewView(v.Model)))

	command.MainThreadHandle[*cmd.ReportErr](v.reportErrCmd)
	return widget
}

func (v *View) reportErrCmd(cmd *cmd.ReportErr) {
	reporterr.Show(v.ViewRoot(), cmd)
}
