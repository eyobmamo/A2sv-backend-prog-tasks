package main

import (
	"gowhere/controllers"
	"gowhere/services"
)

var libraryService services.LibraryService;



func main() {
	libraryService = services.new_library_starter()
	controllers.LibraryController.Console_control()

}
