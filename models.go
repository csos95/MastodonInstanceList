package main

import "time"

// Stats for an instance
type Stats struct {
	DateTime    time.Time `json:"datetime"`
	UserCount   int       `json:"user_count"`
	StatusCount int       `json:"status_count"`
	DomainCount int       `json:"domain_count"`
}

// A Instance of Mastodon
type Instance struct {
	ID           int    `json:"id"`
	URI          string `json:"uri"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Email        string `json:"email"`
	Version      string `json:"version"`
	Stats        Stats  `json:"stats"`
	Thumbnail    string `json:"thumbnail"`
	Topic        string `json:"topic"`
	Note         string `json:"note"`
	Registration string `json:"registration"`
}

type Payload struct {
	Status       string     `json:"status"`
	Instances    []Instance `json:"instances,omitempty"`
	Instance     *Instance  `json:"instance,omitempty"`
	StatsHistory []Stats    `json:"stats_history,omitempty"`
}

// func (i *Instance) String() string {
// 	return fmt.Sprintf("Title: %s\nURI: %s\nDescription: %s\nVersion: %s\nUsers: %d\nStatuses: %d\nConnected Instances: %d\n",
// 		i.Title, i.URI, i.Description, i.Version, i.Stats.UserCount, i.Stats.StatusCount, i.Stats.DomainCount)
// }
