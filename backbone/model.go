package backbone

import (
	"context"
)

type Model interface {
	ModelInit()
	ModelDestroy()
}

type BaseModel struct {
	Ctx       context.Context
	cancelCtx context.CancelFunc
}

func NewBaseModel() *BaseModel {
	return &BaseModel{}
}

func (m *BaseModel) ModelInit() {
	m.Ctx, m.cancelCtx = context.WithCancel(context.Background())
}

func (m *BaseModel) ModelDestroy() {
	m.cancelCtx()
}
