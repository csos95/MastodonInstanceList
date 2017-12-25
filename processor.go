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

func (p *Processor) ReadInstanceStatsHistory(id int) ([]Stats, error) {
	instance, err := p.DB.ReadInstance(id)
	if err != nil {
		return nil, errors.Wrap(err, "[Processor]: failed to read instance")
	}
	statsHistory, err := p.DB.ReadInstanceStatsHistory(instance)
	if err != nil {
		return nil, errors.Wrap(err, "[Processor]: failed to read instance stats history")
	}

	return statsHistory, nil
}

func updateInstance(instance Instance) (Instance, error) {
	response, err := http.Get(fmt.Sprintf("https://%s/api/v1/instance", instance.URI))
	if err != nil {
		return instance, errors.Wrap(err, "failed to get instance stats api")
	}

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&instance)
	if err != nil {
		return instance, errors.Wrap(err, "failed to decode json")
	}
	return instance, nil
}

func compareInstances(a, b Instance) (bool, bool) {
	instanceChanged, statsChanged := false, false
	if a.Title != b.Title || a.Description != b.Description || a.Email != b.Email ||
		a.Version != b.Version || a.Thumbnail != b.Thumbnail {
		instanceChanged = true
	}
	if a.Stats.UserCount != b.Stats.UserCount || a.Stats.StatusCount != b.Stats.StatusCount ||
		a.Stats.DomainCount != b.Stats.DomainCount {
		statsChanged = true
	}
	return instanceChanged, statsChanged
}

func (p *Processor) UpdateInstances() error {
	instances, err := p.DB.ReadInstances()
	if err != nil {
		return errors.Wrap(err, "[Processor]: failed to update instances")
	}
	for _, instance := range instances {
		updatedInstance, err := updateInstance(instance)
		if err != nil {
			log.Println(errors.Wrap(err, "[Processor]: failed to update instance"))
			continue
		}

		instanceChanged, statsChanged := compareInstances(instance, updatedInstance)

		if instanceChanged {
			err = p.DB.UpdateInstance(updatedInstance)
			if err != nil {
				log.Println(errors.Wrap(err, "[Processor]: failed to update instance"))
				continue
			}
		}
		if statsChanged {
			err = p.DB.UpdateStats(updatedInstance)
			if err != nil {
				log.Println(errors.Wrap(err, "[Processor]: failed to update instances"))
			}
		}
	}
	return nil
}
