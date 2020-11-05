package main

import (
    "github.com/AlejandroAM91/orchestra/internal/app/minion"
)

func main() {
    minion := minion.CreateMaster()
    minion.Start()
    minion.Wait()
}

