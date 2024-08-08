package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func Load(config interface{}, name, env, configDir string) error {
	var fileName string
	if env != "" {
		fileName = fmt.Sprintf("%s.%s.json", name, env)
	} else {
		fileName = fmt.Sprintf("%s.json", name)
	}

	filePath := configDir
	filePath = filePath + "/" + fileName
	raw, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(raw, config); err != nil {
		return err
	}

	return nil
}
