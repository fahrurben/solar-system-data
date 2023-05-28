package recover

import (
	"encoding/csv"
	"github.com/fahrurben/solar-system-data/internal/solarsystem/domain/body"
	"github.com/fahrurben/solar-system-data/internal/solarsystem/domain/orbitalparameters"
	"github.com/fahrurben/solar-system-data/internal/solarsystem/domain/physicalcharateristic"
	"github.com/pkg/errors"
	"io"
	"strconv"
)

type ServiceImpl struct {
	bodyRepository     body.Repository
	orbitalRepository  orbitalparameters.Repository
	physicalRepository physicalcharateristic.Repository
}

type SolarModel interface {
	body.Body | orbitalparameters.OrbitalParameters | physicalcharateristic.PhysicalCharacteristic
}

func NewService(bodyRepository body.Repository, orbitalRepository orbitalparameters.Repository, physicalRepository physicalcharateristic.Repository) *ServiceImpl {
	return &ServiceImpl{
		bodyRepository:     bodyRepository,
		orbitalRepository:  orbitalRepository,
		physicalRepository: physicalRepository,
	}
}

func (s *ServiceImpl) Recover(bodyFile io.Reader, orbitalFile io.Reader, physicalFile io.Reader) error {
	arrBody, err := getDataFromFileReader(bodyFile, "Body")
	if err != nil {
		return err
	}

	arrOrbital, err := getDataFromFileReader(orbitalFile, "OrbitalParameters")
	if err != nil {
		return err
	}

	arrPhysical, err := getDataFromFileReader(physicalFile, "PhysicalParameters")
	if err != nil {
		return err
	}

	for _, bodyObj := range arrBody {
		bodyObj, ok := bodyObj.(*body.Body)
		if ok != true {
			return errors.New("Cannot convert object to body")
		}
		existObj, err := s.bodyRepository.Get(bodyObj.Id)
		if existObj != nil {
			continue // Continue when id already exists
		}
		s.bodyRepository.Create(*bodyObj)
		if err != nil {
			return err
		}
	}

	for _, orbitalObj := range arrOrbital {
		orbitalObj, ok := orbitalObj.(*orbitalparameters.OrbitalParameters)
		if ok != true {
			return errors.New("Cannot convert object to orbital parameters")
		}
		existObj, err := s.orbitalRepository.Get(orbitalObj.Id)
		if existObj != nil {
			continue // Continue when id already exists
		}
		s.orbitalRepository.Create(*orbitalObj)
		if err != nil {
			return err
		}
	}

	for _, physicalObj := range arrPhysical {
		physicalObj, ok := physicalObj.(*physicalcharateristic.PhysicalCharacteristic)
		if ok != true {
			return errors.New("Cannot convert object to body")
		}
		existObj, err := s.physicalRepository.Get(physicalObj.Id)
		if existObj != nil {
			continue // Continue when id already exists
		}
		s.physicalRepository.Create(*physicalObj)
		if err != nil {
			return err
		}
	}

	return nil
}

func getDataFromFileReader(csvFile io.Reader, modelType string) ([]any, error) {
	csvReader := csv.NewReader(csvFile)
	var arrData []any

	// Iterate through the body records
	for {
		// Read each record from csv
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		var entity any

		if modelType == "Body" {
			entity, err = getBodyFromCsvRecord(record)
		} else if modelType == "OrbitalParameters" {
			entity, err = getOrbitalFromCsvRecord(record)
		} else if modelType == "PhysicalParameters" {
			entity, err = getPhysicalFromRecord(record)
		} else {
			return nil, errors.New("Not supported model type")
		}

		if err != nil {
			return nil, err
		}
		arrData = append(arrData, entity)

		if err != nil {
			return nil, err
		}
	}

	return arrData, nil
}

func getBodyFromCsvRecord(record []string) (*body.Body, error) {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, err
	}

	typefield := record[1]
	name := record[2]
	description := record[3]

	moons, err := strconv.Atoi(record[4])
	if err != nil {
		return nil, err
	}

	body := &body.Body{
		Id:          id,
		Type:        typefield,
		Name:        name,
		Description: description,
		Moons:       moons,
	}

	return body, nil
}

func getOrbitalFromCsvRecord(record []string) (*orbitalparameters.OrbitalParameters, error) {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, err
	}

	bodyId, err := strconv.Atoi(record[1])
	if err != nil {
		return nil, err
	}

	sideralOrbit, err := strconv.ParseFloat(record[2], 32)
	if err != nil {
		return nil, err
	}

	sideralRotation, err := strconv.ParseFloat(record[3], 32)
	if err != nil {
		return nil, err
	}

	orbital := &orbitalparameters.OrbitalParameters{
		Id:              id,
		BodyId:          bodyId,
		SideralOrbit:    float32(sideralOrbit),
		SideralRotation: float32(sideralRotation),
	}

	return orbital, nil
}

func getPhysicalFromRecord(record []string) (*physicalcharateristic.PhysicalCharacteristic, error) {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, err
	}

	bodyId, err := strconv.Atoi(record[1])
	if err != nil {
		return nil, err
	}

	density, err := strconv.ParseFloat(record[2], 32)
	if err != nil {
		return nil, err
	}

	gravity, err := strconv.ParseFloat(record[3], 32)
	if err != nil {
		return nil, err
	}

	massValue, err := strconv.ParseFloat(record[4], 32)
	if err != nil {
		return nil, err
	}

	massExponent, err := strconv.Atoi(record[5])
	if err != nil {
		return nil, err
	}

	volumeValue, err := strconv.ParseFloat(record[6], 32)
	if err != nil {
		return nil, err
	}

	volumeExponent, err := strconv.ParseFloat(record[7], 32)
	if err != nil {
		return nil, err
	}

	orbital := &physicalcharateristic.PhysicalCharacteristic{
		Id:             id,
		BodyId:         bodyId,
		Density:        float32(density),
		Gravity:        float32(gravity),
		MassValue:      float32(massValue),
		MassExponent:   massExponent,
		VolumeValue:    float32(volumeValue),
		VolumeExponent: float32(volumeExponent),
	}

	return orbital, nil
}
