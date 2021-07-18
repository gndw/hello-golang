package model

func (a *App) Run() (err error) {
	a.Fx.Instance.Run()
	return nil
}