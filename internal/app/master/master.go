package master

import (
    "log"
)

type Master struct {
    sync chan struct{}
}

func CreateMaster() *Master {
  return &Master{
      sync: make(chan struct{}),
  }
}

func (s *Master) Start() {
    log.Println("Starting...")

    go func () {
        log.Println("Started")
        close(s.sync)
        log.Println("Stopped")
    }()
}

func (s Master) Stop() {
    log.Println("Stopping...")
}

func (s Master) Wait() {
    <-s.sync
}
