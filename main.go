package main

import (
	"context"
	"fmt"

	"log"

	"task2/config"
	"task2/db"
	mgtn "task2/migration"
	"task2/utility"
)

func main() {
	mode, path := utility.FlagParse()
	if mode == "" {
		log.Fatal(fmt.Errorf("not correct mode input"))
	}

	cfg, err := config.Init(path)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := db.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := conn.Close(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = mgtn.CheckVersionTable(conn)
	if err != nil {
		log.Fatal(err)
	}

	v, err := mgtn.GetVersion(conn)
	if err != nil {
		log.Fatal(err)
	}

	err = mgtn.Migrate(cfg.Migration.Path, mode, v, conn)
	if err != nil {
		log.Fatal(err)
	}
}
