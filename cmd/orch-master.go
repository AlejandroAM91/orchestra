package main

import (
    "github.com/AlejandroAM91/orchestra/internal/app/master"
)

func main() {
    master := master.CreateMaster()
    master.Start()
    master.Wait()
}

