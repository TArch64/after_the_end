package backbone

import (
	"context"
)

type Model interface {
	ModelInit()
	ModelDestroy()
}

type BaseModel struct {
	*DisposableController
	Ctx       context.Context
	cancelCtx context.CancelFunc
}

func NewBaseModel() *BaseModel {
	return &BaseModel{
		DisposableController: NewDisposableController(),
	}
}

func (m *BaseModel) ModelInit() {
	m.Ctx, m.cancelCtx = context.WithCancel(context.Background())
}

func (m *BaseModel) ModelInitChild(child *BaseModel) {
	child.Ctx, child.cancelCtx = m.Ctx, m.cancelCtx
}

func (m *BaseModel) ModelDestroy() {
	m.cancelCtx()
	m.dispose()
}
