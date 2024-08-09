package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func Load(config interface{}, name, env, configDir string) error {
	fileName := fmt.Sprintf("%s.%s.json", name, env)
	filePath := fmt.Sprintf("%s/%s", configDir, fileName)
	raw, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(raw, config); err != nil {
		return err
	}

	return nil
}
