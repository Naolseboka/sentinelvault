package cfghist

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

type Snapshot struct {
	ID string `json:"id"`
	Version int `json:"version"`
	Config json.RawMessage `json:"config"`
	ActorID string `json:"actor_id"`
	Note string `json:"note,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type Store struct {
	mu sync.Mutex
	snaps []Snapshot
	version int
}

func New() *Store { return &Store{snaps: make([]Snapshot, 0, 32)} }

func (s *Store) Commit(actorID, note string, cfg any) (Snapshot, error) {
	raw, err := json.Marshal(cfg)
	if err != nil {
		return Snapshot{}, err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.version++
	snap := Snapshot{ID: common.NewID(), Version: s.version, Config: raw, ActorID: actorID, Note: note, CreatedAt: common.NowUTC()}
	s.snaps = append(s.snaps, snap)
	return snap, nil
}

func (s *Store) Latest() (Snapshot, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.snaps) == 0 {
		return Snapshot{}, false
	}
	return s.snaps[len(s.snaps)-1], true
}

func (s *Store) Get(version int) (Snapshot, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := len(s.snaps) - 1; i >= 0; i-- {
		if s.snaps[i].Version == version {
			return s.snaps[i], true
		}
	}
	return Snapshot{}, false
}

func (s *Store) List() []Snapshot {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Snapshot, len(s.snaps))
	copy(out, s.snaps)
	return out
}
// maintenance note 78
