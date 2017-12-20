package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type Database struct {
	*sql.DB
}

func OpenDB(cfg Config) (*Database, error) {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s%s",
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBAddress,
		cfg.DBName,
		cfg.DBParameters,
	))
	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}

func (db *Database) CreateTables() error {
	_, err := db.Exec(createInstanceTableSQL)
	if err != nil {
		return errors.Wrap(err, "[Database]: failed to create instance table")
	}
	_, err = db.Exec(createStatsTableSQL)
	if err != nil {
		return errors.Wrap(err, "[Database]: failed to create stats table")
	}
	return nil
}

// 	CREATE

// CreateInstance in the database
func (db *Database) CreateInstance(instance Instance) (int, error) {
	result, err := db.Exec(createInstanceSQL, instance.Title, instance.URI, instance.Description,
		instance.Email, instance.Version, instance.Thumbnail, instance.Topic, instance.Note, instance.Registration)
	if err != nil {
		return -1, errors.Wrap(err, "[Database]: failed to create instance")
	}

	id, _ := result.LastInsertId()
	return int(id), nil
}

// READ

// ReadInstances from the database
func (db *Database) ReadInstances() ([]Instance, error) {
	rows, err := db.Query(readInstancesSQL)
	if err != nil {
		return nil, errors.Wrap(err, "[Database]: failed to read instances from the database")
	}

	instances := make([]Instance, 0)
	for rows.Next() {
		var instance Instance
		err = rows.Scan(&instance.ID, &instance.Title, &instance.URI, &instance.Description,
			&instance.Email, &instance.Version, &instance.Thumbnail, &instance.Topic, &instance.Note, &instance.Registration)
		if err != nil {
			return nil, errors.Wrap(err, "[Database]: failed to scan instances")
		}
		instance.Stats, err = db.ReadInstanceStats(instance)
		if err != nil {
			return nil, errors.Wrap(err, "[Database]: failed to read instance stats")
		}
		instances = append(instances, instance)
	}

	return instances, nil
}

// ReadInstanceStats from the database
func (db *Database) ReadInstanceStats(instance Instance) (Stats, error) {
	var stats Stats
	row := db.QueryRow(readInstanceStatsSQL, instance.ID)
	err := row.Scan(&stats.DateTime, &stats.UserCount, &stats.StatusCount, &stats.DomainCount)
	return stats, errors.Wrap(err, "[Database]: failed to scan instance stats")
}

// UPDATE

// UpdateInstance in the database
func (db *Database) UpdateInstance(instance Instance) error {
	_, err := db.Exec(updateInstanceSQL, instance.Title, instance.Description, instance.Email,
		instance.Version, instance.Thumbnail, instance.Topic, instance.Note, instance.ID, instance.Registration)
	return errors.Wrap(err, "[Database]: failed to update instance")
}

// UpdateStats in the database for an instance
func (db *Database) UpdateStats(instance Instance) error {
	_, err := db.Exec(updateStatsSQL, instance.ID, time.Now(),
		instance.Stats.UserCount, instance.Stats.StatusCount, instance.Stats.DomainCount)
	return errors.Wrap(err, "[Database]: failed to update instance stats")
}

// DELETE
