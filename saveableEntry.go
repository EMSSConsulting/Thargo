package thargo

type Saveable interface {
	Save(destination string) error
}

type SaveableEntry interface {
	Entry
	Saveable
}
