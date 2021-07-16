package builder

type BuilderOptionType int

const (
	Undefined BuilderOptionType = iota
	FxIsProvider
	FxIsStartup
)

type BuilderOption struct {
	optype BuilderOptionType
	list   []interface{}
}