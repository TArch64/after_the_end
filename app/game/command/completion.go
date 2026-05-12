package command

type Waitable interface {
	Wait()
}

type WithCompletion struct {
	competition chan any
}

func NewWithCompletion() *WithCompletion {
	return &WithCompletion{
		competition: make(chan any, 1),
	}
}

func (c *WithCompletion) Complete() {
	c.competition <- struct{}{}
}

func (c *WithCompletion) Wait() {
	<-c.competition
}
