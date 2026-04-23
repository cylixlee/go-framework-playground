//go:build wireinject

package main

import "github.com/google/wire"

func InitializeBat() bat {
	wire.Build(newAnimal, newMouse, newFlyer, newBat)
	return bat{}
}
