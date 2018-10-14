package gatherer

type Factory interface {
	Service() Service
}
