package backbone

type StatefullView[M Model] struct {
	*StatelessView
	Model M
}

func NewStatefullView[M Model](model M) *StatefullView[M] {
	return &StatefullView[M]{
		StatelessView: NewStatelessView(),
		Model:         model,
	}
}

func (b *StatefullView[M]) ViewBeforeInit() {
	b.Model.ModelInit()
}

func (b *StatefullView[M]) ViewDestroy() {
	b.StatelessView.ViewDestroy()
	b.Model.ModelDestroy()
}
