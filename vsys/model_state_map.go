package vsys

type StateMap struct {
	Idx  StateMapIdx
	Data DataEntry
}

func NewStateMap(idx StateMapIdx, data DataEntry) *StateMap {
	return &StateMap{Idx: idx, Data: data}
}

func (sm *StateMap) Serialize() Bytes {
	return append(
		sm.Idx.Serialize(),
		sm.Data.Serialize()...)
}
