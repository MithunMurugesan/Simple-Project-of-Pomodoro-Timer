package main

import (
	"sync"
	"time"
)

type State struct {
	Mode        string `json:"mode"`        // focus or break
	Running     bool   `json:"running"`
	SecondsLeft int    `json:"secondsLeft"`
}

type Session struct {
	State   State
	mu      sync.Mutex
	ticker  *time.Ticker
	stopCh  chan struct{}
}

var session = &Session{
	State: State{
		Mode:        "focus",
		Running:     false,
		SecondsLeft: 25 * 60,
	},
}

func (s *Session) StartTimer() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.State.Running {
		return
	}
	s.State.Running = true
	s.stopCh = make(chan struct{})
	s.ticker = time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-s.ticker.C:
				s.tick()
			case <-s.stopCh:
				return
			}
		}
	}()
}

func (s *Session) StopTimer() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.State.Running {
		return
	}
	s.State.Running = false
	s.ticker.Stop()
	close(s.stopCh)
}

func (s *Session) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.State.Running {
		s.ticker.Stop()
		close(s.stopCh)
	}
	s.State.Mode = "focus"
	s.State.Running = false
	s.State.SecondsLeft = 25 * 60
}

func (s *Session) tick() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.State.SecondsLeft > 0 {
		s.State.SecondsLeft--
	} else {
		if s.State.Mode == "focus" {
			s.State.Mode = "break"
			s.State.SecondsLeft = 5 * 60
		} else {
			s.State.Mode = "focus"
			s.State.SecondsLeft = 25 * 60
		}
	}
}
