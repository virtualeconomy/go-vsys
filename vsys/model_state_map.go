package vsys

// StateMap is struct for representing state map of the contracts.
type StateMap struct {
	Idx  StateMapIdx
	Data DataEntry
}

// NewStateMap constructs state map from given idx and data entry.
func NewStateMap(idx StateMapIdx, data DataEntry) *StateMap {
	return &StateMap{Idx: idx, Data: data}
}

// Serialize serializes StateMap to Bytes.
func (sm *StateMap) Serialize() Bytes {
	return append(
		sm.Idx.Serialize(),
		sm.Data.Serialize()...)
}
