package event

type EventKind string
type ImplementationType string

const (

	// MergeToMaster is whenever a pull request is merged to the
	// master branch
	MergeToMaster EventKind = "Merge To Master"
)

type Event interface {
	Kind() EventKind
	Type() ImplementationType
}

type Broker interface {
	ConcurrentWatch(queue *Queue) chan error
}
