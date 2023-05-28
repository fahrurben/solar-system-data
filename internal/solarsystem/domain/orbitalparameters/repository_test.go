package orbitalparameters

import (
	"fmt"
	"github.com/fahrurben/solar-system-data/internal/solarsystem/platform/db"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// LoadConfig reads configuration f`rom file or environment variables.
func LoadConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("test")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
	}
}

var repository *RepositoryImpl
var mysql *sqlx.DB

func init() {
	LoadConfig("../../../../configs/")
	var err error
	mysql, err = db.New(viper.GetString("DATABASE_URL"))
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
	}
	repository = NewRepository(mysql)
}

func truncateDatabase() error {
	_, err := mysql.Exec("truncate table body")
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
	}

	_, err = mysql.Exec("truncate table physical_data")
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
	}

	_, err = mysql.Exec("truncate table orbital_parameters")
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
	}

	return err
}

func TestServiceImpl_CreateOrbitalParameters(t *testing.T) {
	err := truncateDatabase()
	if err != nil {
		assert.Fail(t, "Truncate database failed")
	}

	orbitalParameters := OrbitalParameters{
		Id:              1,
		BodyId:          1,
		SideralOrbit:    1,
		SideralRotation: 1,
	}

	_, err = repository.Create(orbitalParameters)
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
		assert.Fail(t, "Create Orbital Parameter failed")
	}
}
