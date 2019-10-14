package magic

// SrcProvider can be used to deliver input into a magic.Manager
type SrcProvider interface {
	Scan() bool
	Text() string
}
