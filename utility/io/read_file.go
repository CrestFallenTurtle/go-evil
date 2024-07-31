package io

import (
	"io/ioutil"

	"github.com/CrestFallenTurtle/go-evil/utility/structure/json"
	notify "github.com/CrestFallenTurtle/notify_handler"
)

// Reads the contents of the file found in the json data structure
func Read_file(data_object *json.Json_t) {

	content, err := ioutil.ReadFile(data_object.File_path)

	if err != nil {
		notify.Error(err.Error(), "io.Read_file()", 1)
	}

	data_object.File_gut = string(content)

}
