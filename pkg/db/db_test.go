package db

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/floriwan/srcm/pkg/config"
	"github.com/floriwan/srcm/pkg/db/model"
)

type tc struct {
	u model.User
	e error
}

var dbCreateUserTests = []struct {
	u model.User
	e error
}{
	{
		u: model.User{},
		e: fmt.Errorf("NOT NULL constraint failed: users.password"),
	},
	{
		u: model.User{Password: "secret"},
		e: fmt.Errorf("NOT NULL constraint failed: users.email"),
	},
	{
		u: model.User{Password: "secret", Email: "kermit@sesamstrasse.de"},
	},
	{
		u: model.User{Password: "secret", Email: "kermit@sesamstrasse.de"},
		e: fmt.Errorf("UNIQUE constraint failed: users.email"),
	},
	{
		u: model.User{Password: "secret", Email: "grobi@sesamstrasse.de"},
	},
}

func TestMain(m *testing.M) {
	config.GlobalConfig = config.Config{}
	testResult := m.Run()
	os.Exit(testResult)
}

func TestCreateDifferentUser(t *testing.T) {
	config.GlobalConfig.DbSqliteFilename = filepath.Join(t.TempDir(), "test.db")
	Initialize()
	Migrate()

	for _, tc := range dbCreateUserTests {
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

		if tc.u.Active != true {
			t.Fatalf("user '%v' should be active, but is '%v'\n", tc.u.Email, tc.u.Active)
		}

		fmt.Printf("create user id:%v\n", tc.u.ID)
	}

	users := []model.User{}
	r := Instance.Find(&users)
	if r.RowsAffected != 2 {
		t.Fatalf("should get 2 rows, but got %v\n", r.RowsAffected)
	}

	fmt.Printf("\n%+v\n", users)

}
