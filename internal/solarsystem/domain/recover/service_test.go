package recover

import (
	"fmt"
	"github.com/fahrurben/solar-system-data/internal/solarsystem/domain/body"
	"github.com/fahrurben/solar-system-data/internal/solarsystem/domain/orbitalparameters"
	"github.com/fahrurben/solar-system-data/internal/solarsystem/domain/physicalcharateristic"
	"github.com/fahrurben/solar-system-data/internal/solarsystem/platform/db"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
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

var bodyRepository *body.RepositoryImpl
var orbitalRepository *orbitalparameters.RepositoryImpl
var physicalRepository *physicalcharateristic.RepositoryImpl
var mysql *sqlx.DB

func init() {
	LoadConfig("../../../../configs/")
	var err error
	mysql, err = db.New(viper.GetString("DATABASE_URL"))
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
	}
	bodyRepository = body.NewRepository(mysql)
	orbitalRepository = orbitalparameters.NewRepository(mysql)
	physicalRepository = physicalcharateristic.NewRepository(mysql)
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

func TestRecover(t *testing.T) {
	truncateDatabase()

	serviceImpl := NewService(bodyRepository, orbitalRepository, physicalRepository)

	bodyFile, err := os.Open("./testfiles/body.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	orbitalFile, err := os.Open("./testfiles/orbital.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	physicalFile, err := os.Open("./testfiles/physical.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	err = serviceImpl.Recover(bodyFile, orbitalFile, physicalFile)
	assert.Nil(t, err)
}
