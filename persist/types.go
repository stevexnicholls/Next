package persist

// Value returned from a persistence medium
type Value struct {
	Value []byte `json:"value"`
}
