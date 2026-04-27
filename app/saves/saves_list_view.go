package saves

import (
	"after_the_end/backbone"
	"after_the_end/db/model"

	qt "github.com/mappu/miqt/qt6"
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
	widget.SetProperty("bg-reset", qt.NewQVariant8(true))

	column := qt.NewQVBoxLayout(widget)
	column.SetContentsMargins(0, 0, 0, 0)

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
