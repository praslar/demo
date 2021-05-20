
package utils

import (

	"github.com/praslar/lib/response"
	"github.com/sirupsen/logrus"

	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

type (
	//Status format from status pkg
	Status = response.ResponseStatus

	GenStatus struct {
		Success      Status
		BadRequest   Status 
		Unauthorized Status
		Internal     Status
		Database     Status
	}

	statuses struct {
		Gen              GenStatus
	}
)

var (
	all  *statuses
	once sync.Once
)

// Init load statuses from the given config file.
// Init panics if cannot access or error while parsing the config file.
func Init(conf string) {
	once.Do(func() {
		f, err := os.Open(conf)
		if err != nil {
			logrus.Errorf("Fail to open status file, %v", err)
			panic(err)
		}
		all = &statuses{}
		if err := yaml.NewDecoder(f).Decode(all); err != nil {
			logrus.Errorf("Fail to parse status file data to statuses struct, %v", err)
			panic(err)
		}
	})
}

// all return all registered statuses.
// all will load statuses from configs/Status.yml if the statuses has not initialized yet.
func load(err string) *statuses {
	conf := os.Getenv("STATUS_PATH")
	if conf == "" {
		conf = "conf/status.yml"
	}
	Init(conf)
	if err != "" {
		all.Gen.BadRequest.XMessage = err
	}
	return all
}

func Gen(err string) GenStatus {
	return load(err).Gen
}
