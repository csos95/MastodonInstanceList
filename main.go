package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func getInstance(uri string) (Instance, error) {
	var instance Instance

	response, err := http.Get(fmt.Sprintf("https://%s/api/v1/instance", uri))
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

func preload(config Config) error {
	db, err := OpenDB(config)
	if err != nil {
		return err
	}

	err = db.CreateTables()
	if err != nil {
		return err
	}

	// load some starter instances
	// instanceURIs := []string{
	// 	"a.weirder.earth",
	// 	"bookwitty.social",
	// 	"tootcn.com",
	// 	"i.write.codethat.sucks",
	// 	"mamot.fr",
	// 	"mastodon.tetaneutral.net",
	// 	"mstdn.fr",
	// 	"social.alex73630.xyz",
	// 	"social.infranix.eu",
	// 	"social.taker.fr",
	// }
	// instanceTopics := []string{"be weird", "book lovers", "Chinese", "Code", "FL", "Fr/GP", "Fr/GP", "Fr/GP", "Fr/GP", "Fr/GP"}
	// instanceNotes := []string{"silence instances", "Fr/Eng", "", "", "", "", "", "", "Fr/Eng/fet w/pawoo", ""}
	// instanceRegistrations := []string{"open", "open", "open", "open", "open", "open", "open", "open", "open", "open"}
	instanceURIs := []string{"mastodon.social", "toot.cafe"}
	instanceTopics := []string{"general", "Code"}
	instanceNotes := []string{"the largest mastodon instance", ""}
	instanceRegistrations := []string{"open", "open"}
	for i, uri := range instanceURIs {
		instance, err := getInstance(uri)
		if err != nil {
			log.Println(err)
		}
		instance.Topic = instanceTopics[i]
		instance.Note = instanceNotes[i]
		instance.Registration = instanceRegistrations[i]
		instance.ID, err = db.CreateInstance(instance)
		if err != nil {
			log.Println(err)
		}

		err = db.UpdateStats(instance)
		if err != nil {
			log.Println(err)
		}
	}
	db.Close()

	fmt.Println("done preloading")
	return nil
}

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	// err = preload(config)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// run the server
	server, err := NewServer(config)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Run()
	if err != nil {
		log.Println(err)
	}
}
