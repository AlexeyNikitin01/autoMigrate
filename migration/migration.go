package migration

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	
	"github.com/jackc/pgx/v5"
)

func CheckVersionTable(conn *pgx.Conn) error {
	var count int
	err := conn.QueryRow(context.TODO(), "SELECT COUNT(*) FROM information_schema.tables where table_name='go_migrations';").Scan(&count)
	if err != nil {
		return err
	}
	if !(count == 1) {
		err := createVersionTable(conn)
		if err != nil {
			return err
		}
		log.Println("create table go_microservice")
	}
	return nil
}

func createVersionTable(conn *pgx.Conn) error {
	_, err := conn.Exec(context.TODO(), "CREATE TABLE IF NOT EXISTS go_migrations (id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY, version integer NOT NULL, created_at timestamp DEFAULT now() NOT NULL);")
	if err != nil {
		return fmt.Errorf("don`t create version table")
	}
	return nil
}

func Migrate(path, flg string, v int, conn *pgx.Conn) error {
	switch {
	case flg == "up":
		for i := v+1; i <= 2; i++ {
			content, err := readFileMigration(i, "up", path)
			if err != nil {
				return err
			}
			err = insertVersion(conn, i)
			if err != nil {
				return err
			}
			err = execMigration(content, conn)
			if err != nil {
				return err
			}
		}
	case flg == "down" && v > 0:
		content, err := readFileMigration(v, "down", path)
		if err != nil {
			return err
		}
		err = insertVersion(conn, v-1)
		if err != nil {
			return err
		}
		err = execMigration(content, conn)
		if err != nil {
			return err
		}
	case flg == "reset":
		err := dropVersion(conn)
		if err != nil {
			return err
		}
	case flg == "version":
		fmt.Println(v)
	}
	return nil
}

func readFileMigration(version int, flag, path string) (string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return "", fmt.Errorf("no find migration direcroty")
	}
	for _, file := range files {
		if strings.Split(file.Name(), "_")[0] == strconv.Itoa(version) {
			content, err := os.ReadFile(path + "/" + file.Name())
			if err != nil {
				return "", fmt.Errorf("no correct file %s", file)
			}
			if strings.Split(file.Name(), ".")[1] == flag {
				return string(content), nil
			}
		}
	}
	return "", fmt.Errorf("no find migration")
}

func insertVersion(conn *pgx.Conn, v int) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO go_migrations (version) VALUES($1)", v)
	if err != nil {
		return fmt.Errorf("not insert version")
	}
	return nil
}

func execMigration(content string, conn *pgx.Conn) error {
	_, err := conn.Exec(context.TODO(), content)
	if err != nil {
		return fmt.Errorf("not exec migration")
	}
	return nil
}

func dropVersion(conn *pgx.Conn) error {
	_, err := conn.Exec(context.TODO(), "DROP TABLE IF EXISTS go_migrations; DROP TABLE IF EXISTS widgets;")
	if err != nil {
		return fmt.Errorf("not drop version")
	}
	return err
}

func GetVersion(conn *pgx.Conn) (int, error) {
	v := 0
	row := false
	err := conn.QueryRow(context.TODO(), "SELECT EXISTS (SELECT version FROM go_migrations ORDER BY id DESC)").Scan(&row)
	if err != nil {
		return 0, fmt.Errorf("uncorrect query get version")
	}
	if !row {
		return 0, nil
	}
	err = conn.QueryRow(context.TODO(), "SELECT version FROM go_migrations ORDER BY id DESC").Scan(&v)
	if err != nil {
		return 0, fmt.Errorf("uncorrect query get version")
	}
	return v, nil
}
