package saves

import (
	"after_the_end/backbone"
	"after_the_end/backbone/styled"
	"after_the_end/db/model"

	"github.com/mappu/miqt/qt"
)

type ListView struct {
	*backbone.StatelessView
	gameSaves []*model.GameSave
	onDelete  func(gameSave *model.GameSave)
}

type ListViewOptions struct {
	GameSaves []*model.GameSave
	OnDelete  func(gameSave *model.GameSave)
}

func NewListView(options *ListViewOptions) *ListView {
	return &ListView{
		StatelessView: backbone.NewStatelessView(),
		gameSaves:     options.GameSaves,
		onDelete:      options.OnDelete,
	}
}

func (v *ListView) ViewInit() *qt.QWidget {
	widget := qt.NewQWidget2()
	widget.SetObjectName("saves_list")
	widget.SetStyleSheet(styled.S("#saves_list", styled.Transparent))

	column := qt.NewQVBoxLayout(widget)
	for _, gameSave := range v.gameSaves {
		view := NewSaveView(&SaveViewOptions{
			GameSave: gameSave,

			OnDeleted: func() {
				v.onDelete(gameSave)
			},
		})

		column.AddWidget(v.Mount(view))
	}

	return widget
}
