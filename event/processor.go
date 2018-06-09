package event

type Processor interface {
	Process(target, name string, data interface{}) (processed bool, err error)
}
