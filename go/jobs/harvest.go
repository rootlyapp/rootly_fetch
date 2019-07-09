package jobs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"samfaina/dev/rootly_fetch/config"
	"samfaina/dev/rootly_fetch/database"
	"samfaina/dev/rootly_fetch/models"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	conf = config.New()
}

// FetchPlants to grab each plant from Trefle API
func FetchPlants() {
	idList := database.GetPlantsIDList()

	// get database connection
	client, ctx, ctxCancel := database.GetDatabaseConnection()
	for i := 0; i < len(idList); i++ {
		plant := fetchPlant(idList[i])
		if plant.MainSpecies.CompleteData {
			database.SavePlant(ctx, client, plant)
		}

		if conf.DebugMode {
			perc := (float64(i+1) / float64(len(idList))) * 100
			log.Printf("Progress: %.3f %%   [%v/%v]", perc, i+1, len(idList))
		}
	}

	if conf.DebugMode {
		log.Println("Yay! All plants saved successfully!")
	}

	ctxCancel()
}

func fetchPlant(plantID int) models.Plant {
	url := fmt.Sprintf(conf.Trefle.UrlPlant, plantID, conf.Trefle.Token)
	response, err := http.Get(url)
	plant := models.Plant{}
	if err != nil {
		log.Fatalf("The HTTP request failed with error %s\n", err)
	} else if response.StatusCode != http.StatusOK {
		data, _ := ioutil.ReadAll(response.Body)
		bodyString := string(data)
		log.Fatalf("[%v] %v", response.StatusCode, bodyString)

	} else {
		data, _ := ioutil.ReadAll(response.Body)
		dataBytes := []byte(data)

		err := json.Unmarshal(dataBytes, &plant)
		if err != nil {
			log.Fatal(err)
		}
		return plant
	}

	return plant
}
