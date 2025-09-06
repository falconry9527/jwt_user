package config

import (
	"path"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
)

func init() {
	// 自动执行，不用调用
	currentAbPath := getCurrentAbPathByCaller()
	tomlFile, err := filepath.Abs(currentAbPath + "/config.toml")
	//tomlFile, err := filepath.Abs(currentAbPath + "/configV22.toml")
	if err != nil {
		panic("read toml file err: " + err.Error())
		return
	}
	if _, err := toml.DecodeFile(tomlFile, &Config); err != nil {
		panic("read toml file err: " + err.Error())
		return
	}
}

func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
