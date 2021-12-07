package main

import (
	"foreigner_guide/src/controllers"
	"foreigner_guide/src/database/migration"
)


func main() {
	migration.Init()
  controllers.Init()		
}