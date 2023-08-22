package seeds

import (
	"fmt"
)

type Seeder interface{
	Run() error
}

func DBSeed() error {
	fmt.Println("Running seed...")

	seeders := []Seeder{
		seedUsers{},
		seedCategories{},
	}

	for _, seeder := range seeders {
		err := seeder.Run()

		if err != nil {
			return err
		}
	}

	return nil
	
}