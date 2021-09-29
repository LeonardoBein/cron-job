package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/LeonardoBein/cron-job/entity"
	"github.com/LeonardoBein/cron-job/lang"
)

func GetConfig(pathFile string) []entity.ConfigCronJob {

	var allConfig []entity.ConfigCronJob

	_, err := os.Stat(pathFile)

	if err != nil {
		panic("[" + pathFile + "]" + lang.ErrConfigPathNotFound)
	}

	filepath.Walk(pathFile, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		jsonFile, err := os.Open(path)

		if err != nil {
			return err
		}

		bytesValue, _ := ioutil.ReadAll(jsonFile)

		var config entity.ConfigCronJob

		err = json.Unmarshal(bytesValue, &config)

		if err != nil {
			return err
		}

		allConfig = append(allConfig, config)

		return nil
	})

	return allConfig

}
