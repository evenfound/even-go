package repo

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

	//"github.com/evenfound/even-go/repo/migrations"
)

type Migration interface {
	Up(repoPath string, dbPassword string, testnet bool) error
	Down(repoPath string, dbPassword string, testnet bool) error
}

var (
	ErrUnknownSchema = errors.New("unable to migrate unknown schema")

	// TODO: commented mode refactoring

	Migrations = [17]Migration{} //{
		//Migration{},
		//Migration{},
		//Migration{},
		//Migration{},
		//Migration{},
		//Migration{},
		//Migration{},
		//Migration{},
		//Migration{},
		//Migration{},
		//Migration{},
		//Migration{},
		//Migration{},
		//Migration{},
		//Migration{},
		//Migration{},
		//Migration{},
	//}
)

// MigrateUp looks at the currently active migration version
// and will migrate all the way up (applying all up migrations).
func MigrateUp(repoPath, dbPassword string, testnet bool) error {
	version, err := ioutil.ReadFile(path.Join(repoPath, "repover"))
	if err != nil && !os.IsNotExist(err) {
		return err
	} else if err != nil && os.IsNotExist(err) {
		log.Noticef("missing repo version file, migrating from 0")
		version = []byte("0")
	}
	v, err := strconv.Atoi(strings.Trim(string(version), "\n"))
	if err != nil {
		return err
	}
	if v > len(Migrations) {
		log.Errorf("binary can migrate schemas up to version %03d but this schema is already at %03d", len(Migrations), v)
		return ErrUnknownSchema
	}
	x := v
	for _, m := range Migrations[v:] {
		log.Noticef("running migration %03d changing schema to version %03d...\n", x, x+1)
		err := m.Up(repoPath, dbPassword, testnet)
		if err != nil {
			log.Error(err)
			return err
		}
		x++
	}
	return nil
}
