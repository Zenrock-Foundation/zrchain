package main

import (
	"sync"
)

// ValidatorSyncState provides thread-safe access to validator state
// during epoch transitions to prevent stale reads.
type ValidatorSyncState struct {
	mu      sync.RWMutex
	epoch   uint64
	syncing bool
}

// NewValidatorSyncState creates a new sync state tracker.
func NewValidatorSyncState() *ValidatorSyncState {
	return &ValidatorSyncState{}
}

// BeginSync marks the start of a synchronization cycle.
func (s *ValidatorSyncState) BeginSync(epoch uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.epoch = epoch
	s.syncing = true
}

// EndSync marks the completion of a synchronization cycle.
func (s *ValidatorSyncState) EndSync() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.syncing = false
}

// IsSyncing returns whether a sync operation is in progress.
func (s *ValidatorSyncState) IsSyncing() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.syncing
}

// CurrentEpoch returns the current epoch being synced.
func (s *ValidatorSyncState) CurrentEpoch() uint64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.epoch
}
