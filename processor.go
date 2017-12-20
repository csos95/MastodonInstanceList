package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type Processor struct {
	DB *Database
}

func NewProcessor(config Config) (*Processor, error) {
	db, err := OpenDB(config)
	if err != nil {
		return nil, errors.Wrap(err, "[Processor]: failed to create new processor")
	}
	return &Processor{db}, nil
}

func (p *Processor) ReadInstances() ([]Instance, error) {
	instances, err := p.DB.ReadInstances()
	if err != nil {
		return nil, errors.Wrap(err, "[Processor]: failed to read instances")
	}
	return instances, nil
}

func updateInstance(instance *Instance) error {
	response, err := http.Get(fmt.Sprintf("https://%s/api/v1/instance", instance.URI))
	if err != nil {
		return errors.Wrap(err, "failed to get instance stats api")
	}

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&instance)
	if err != nil {
		return errors.Wrap(err, "failed to decode json")
	}
	return nil
}

func (p *Processor) UpdateInstances() error {
	instances, err := p.DB.ReadInstances()
	if err != nil {
		return errors.Wrap(err, "[Processor]: failed to update instances")
	}
	for _, instance := range instances {
		err = updateInstance(&instance)
		if err != nil {
			log.Println(errors.Wrap(err, "[Processor]: failed to update instance"))
			continue
		}
		err = p.DB.UpdateInstance(instance)
		if err != nil {
			log.Println(errors.Wrap(err, "[Processor]: failed to update instance"))
			continue
		}
		err = p.DB.UpdateStats(instance)
		if err != nil {
			log.Println(errors.Wrap(err, "[Processor]: failed to update instances"))
		}
	}
	return nil
}
