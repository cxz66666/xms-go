package main

import (
	"xms/models"
	"xms/pkg/logging"
	"xms/pkg/setting"
)

// @title xms
// @license.name MIT
// @version 1.0
// @host xms.zjueva.net

func main() {
	// packages init, you can also use `init` function to init package one by one, but
	// init function will be called in order of dependency, so much time it's not very obviously
	// so we rename `init` to `Setup` and call them in our needed orders
	setting.Setup()

	models.Setup()
	logging.Setup()
}
