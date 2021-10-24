package main

type Cipher int

const (
	Ordinal Cipher = iota
	OrdinalReverse
	FullReduction
	FullReductionReverse
)

func (c Cipher) String() string {
	switch c {
	case Ordinal:
		return "Ordinal"
	case OrdinalReverse:
		return "Reverse Ordinal"
	case FullReduction:
		return "Full Reduction"
	case FullReductionReverse:
		return "Reverse Full Reduction"
	default:
		return "Unknown"
	}
}

type Word struct {
	originalString string
	wordDisplay    string
	valueDisplay   string
}

