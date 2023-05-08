package domain

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
	LoadConfig("../../../configs/")
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

	return err
}

func TestCreateBody(t *testing.T) {
	err := truncateDatabase()
	if err != nil {
		assert.Fail(t, "Truncate database failed")
	}

	body := Body{
		Id:          1,
		Type:        "test",
		Name:        "test",
		Description: "test",
		Moons:       1,
	}

	_, err = repository.CreateBody(body)
	if err != nil {
		assert.Fail(t, "Create body failed")
	}
}

func TestCreatePhysicalCharacteristic(t *testing.T) {
	physicalCharacteristic := PhysicalCharacteristic{
		Id:             1,
		BodyId:         1,
		Density:        1,
		Gravity:        1,
		MassValue:      1,
		MassExponent:   1,
		VolumeValue:    1,
		VolumeExponent: 1,
	}

	_, err := repository.CreatePhysicalCharacteristic(physicalCharacteristic)
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
		assert.Fail(t, "Create Physical Characteristic failed")
	}
}
