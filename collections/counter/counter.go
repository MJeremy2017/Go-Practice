package counter

// unexported identifier
type alertCounter int

// exported identifier
func New(n int) alertCounter{
	return alertCounter(n)
}
