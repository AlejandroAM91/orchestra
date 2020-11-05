package minion

import (
    "log"
)

type Minion struct {
    sync chan struct{}
}

func CreateMinion() *Minion {
  return &Minion{
      sync: make(chan struct{}),
  }
}

func (s *Minion) Start() {
    log.Println("Starting...")

    go func () {
        log.Println("Started")
        close(s.sync)
        log.Println("Stopped")
    }()
}

func (s Minion) Stop() {
    log.Println("Stopping...")
}

func (s Minion) Wait() {
    <-s.sync
}
