package lib

import (
	"encoding/json"
	"errors"
	"os"

  "github.com/adrg/xdg"
)

func getConfigPath(config_path string) string {
  if FileExists(config_path) {
    return config_path
  } else {
    configFilePath, _ := xdg.ConfigFile("mdrss/config.json")
    return configFilePath
  }
}

func FileExists(filename string) bool {
  if _, err := os.Stat(filename); err != nil {
    return false
  }
  return true
}

func ReadConfig(config_path string) (Config, error) {
  var config Config
  configPath := getConfigPath(config_path)
  if FileExists(configPath) {
    configContent, _ := os.ReadFile(configPath)
    jsonErr := json.Unmarshal(configContent, &config)
    if jsonErr != nil {
      return config, errors.New("Error when reading config file.")
    }
    return config, nil
  }
  configFilePath, _ := xdg.ConfigFile("mdrss/config.json")
  return config, errors.New("Config file not found. Please add it at " + configFilePath)
}
