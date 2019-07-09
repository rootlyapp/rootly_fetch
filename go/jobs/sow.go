package jobs

import (
	"log"

	"github.com/joho/godotenv"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"samfaina/dev/rootly_fetch/config"
	"strconv"
)

// PageSize Max rows per page
const PageSize = "1000"

// TotalPages Total available pages
const TotalPages = 148

// conf App environment config
var conf *config.Config

type plantID struct {
	ID int `json:"id"`
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	conf = config.New()

}

// FetchPlantsIds from trefle api
func FetchPlantsIds() []int {
	plantsIdsList := make([]int, 0, 0)

	url := fmt.Sprintf(conf.Trefle.UrlIds, conf.Trefle.Token, PageSize)
	for i := 1; i <= TotalPages; i++ {

		response, err := http.Get(url + strconv.Itoa(i))
		if err != nil {
			log.Fatalf("The HTTP request failed with error %s\n", err)
		} else if response.StatusCode != http.StatusOK {
			data, _ := ioutil.ReadAll(response.Body)
			bodyString := string(data)
			log.Fatalf("[%v] %v", response.StatusCode, bodyString)

		} else {
			data, _ := ioutil.ReadAll(response.Body)
			dataBytes := []byte(data)

			// bodyString := string(data)
			// fmt.Printf("[%v] %v", response.StatusCode, bodyString)

			ids := []plantID{}
			err := json.Unmarshal(dataBytes, &ids)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			// fmt.Println(plant1)
			length := len(ids)
			slice := make([]int, 0, 0)

			for j := 0; j < length; j++ {
				slice = append(slice, ids[j].ID)
			}

			if conf.DebugMode {
				log.Printf("Page %v/%v - Fetched: %v\n", i, TotalPages, len(plantsIdsList))
			}

			plantsIdsList = append(plantsIdsList, slice...)
		}
	}

	return plantsIdsList
}
