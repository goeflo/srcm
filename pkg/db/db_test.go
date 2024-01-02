package db

import (
	"fmt"
	"os"
	"testing"

	"github.com/floriwan/srcm/pkg/config"
	"github.com/floriwan/srcm/pkg/db/model"
)

type tc struct {
	u model.User
	e error
}

func TestMain(m *testing.M) {
	config.GlobalConfig = config.Config{}
	testResult := m.Run()
	os.Exit(testResult)
}

func TestCreateDifferentUser(t *testing.T) {
	config.GlobalConfig.DbSqliteFilename = t.TempDir() + "test.db"
	Initialize()
	Migrate()

	tcs := []tc{
		{
			u: model.User{},
			e: fmt.Errorf("NOT NULL constraint failed: users.email"),
		},
		{
			u: model.User{Email: "kermit@sesamstrasse.de"},
		},
		{
			u: model.User{Email: "kermit@sesamstrasse.de"},
			e: fmt.Errorf("UNIQUE constraint failed: users.email"),
		},
		{
			u: model.User{Email: "grobi@sesamstrasse.de"},
		},
	}

	for _, tc := range tcs {
		r := Instance.Debug().Create(&tc.u)
		if tc.e != nil {
			if r.Error.Error() != tc.e.Error() {
				t.Fatalf("expected '%v', but got '%v'\n", tc.e, r.Error)
			}
		} else {
			if r.Error != nil {
				t.Fatalf("exptected no error, but got '%v'\n", r.Error)
			}
		}
		fmt.Printf("create user id:%v\n", tc.u.ID)
	}

}
