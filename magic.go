package magic

// Manager is the main interaction point with the magic package. It coordinates
// all underlying functions and calls to external services
type Manager struct {
	cursor *cursor
}

// NewManager initializes a new manager instance
func NewManager(src SrcProvider) *Manager {
	return &Manager{
		cursor: newCursor(src),
	}
}
