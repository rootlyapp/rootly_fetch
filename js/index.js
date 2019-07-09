const axios = require("axios");
const fs = require("fs");
const db = require("./db");

const TOKEN = "M2hRU3hyOFYwZERHUjlYNzB5SCtnQT09";
const URL = "http://trefle.io/api/plants/";
const TOTAL_PAGES = 4906;

let idList = require("./plantsID.json") || [];
const args = process.argv.slice(2)[0];
// let plantList = [];

// axios interceptor
axios.interceptors.response.use(undefined, function axiosRetryInterceptor(err) {
  var config = err.config;
  // If config does not exist or the retry option is not set, reject
  if (!config || !config.retry) {
    console.error("Error intercepted!: ", err);
    return Promise.reject(err);
  }

  // Set the variable for keeping track of the retry count
  config.__retryCount = config.__retryCount || 0;

  // Check if we've maxed out the total number of retries
  if (config.__retryCount >= config.retry) {
    // Reject with the error
    return Promise.reject(err);
  }

  // Increase the retry count
  config.__retryCount += 1;

  // Create new promise to handle exponential backoff
  var backoff = new Promise(function(resolve) {
    setTimeout(function() {
      resolve();
    }, config.retryDelay || 1);
  });

  // Return the promise in which recalls axios to retry the request
  return backoff.then(function() {
    console.log("Retrying 'cause ECONNRESET");
    return axios(config);
  });
});

const sleep = ms => new Promise(resolve => setTimeout(resolve, ms));

const getAllPlantID = async () => {
  console.log("FETCHING ALL PLANTS ID\n");
  let list = [];
  try {
    for (let i = 1; i <= TOTAL_PAGES; i++) {
      //   if (i % 1000 === 0) {
      //     console.log("paused");
      //     await sleep(5000);
      //     console.log("resumed");
      //   }
      const resp = await axios.get(URL, {
        params: {
          token: TOKEN,
          page: i,
          retry: 5,
          retryDelay: 2000
        }
      });
      const items = resp.data.map(it => it.id);
      //   idList = idList.concat(items);
      list = list.concat(items);
      console.log("idList length: %i    |   page: %i", list.length, i);
    }
    console.log("All plant ID fetched\n\n");
    await db.saveIdCollection(list);

    return Promise.resolve(list);
  } catch (error) {
    console.error(error);
    return Promise.reject(error);
  }
};

const getAllPlants = async source => {
  let _source = idList;
  if (source) {
    _source = source;
  }
  const initialPlant = 114787;
  console.log("FETCHING ALL PLANTS DATA\n");
  const _length = _source.length || 0;
  let list = [];
  try {
    await db.loadDatabase();
    let counter = 0;
    for (let i = initialPlant; i < _length; i++) {
      const resp = await axios.get(`${URL}${_source[i]}`, {
        params: {
          token: TOKEN
        },
        retry: 5,
        retryDelay: 2000
      });
      list.push(resp.data);

      console.log(
        "Status: %i/%i  ============  %s%",
        i + 1,
        _length,
        (((i + 1) / _length) * 100).toFixed(3)
      );
      if(resp.data.main_species.complete_data){
        await db.addPlantToCollection(resp.data);
        ++counter;
        console.log("Completed plants status: %i/i% [%s%]", counter, _length+1, (((counter) / _length+1) * 100).toFixed(3) );
        
      }
    }
    console.log("All plants fetched!\n\n");
    return Promise.resolve(list);
  } catch (error) {
    console.error(error);
    return Promise.reject(error);
  }
};

const writeFile = (data, filename) => {
  console.log("WRITING DATA TO JSON\n");
  const _data = JSON.stringify(data);
  return new Promise(resolve, reject => {
    fs.writeFile(filename, _data, "utf8", err => {
      if (err) {
        console.log("An error occured while writing JSON Object to File.");
        reject(err);
      } else {
        console.log("JSON file has been saved.");
        resolve(data);
      }
    });
  });
};

// fetch idList
async function fetchPlantIdList() {
  try {
    const ids = await getAllPlantID();
    await writeFile(data, "plantsID.json");
    return Promise.resolve(ids);
  } catch (error) {
    console.log("TCL: fetchPlantIdList -> error", error);
    return Promise.reject(error);
  } finally {
    console.timeEnd("performance");
  }
}

// fetch plant list from idFile
async function fetchPlantListFromFile() {
  try {
    const plants = await getAllPlants();
    await writeFile(plants, "plants.json");

    return Promise.resolve(plants);
  } catch (error) {
    console.log("TCL: fetchPlantListFromFile -> error", error);
    return Promise.reject(error);
  } finally {
    console.timeEnd("performance");
  }
}
// fetch all
async function fetchAllDataToFile() {
  try {
    const ids = await getAllPlantID();
    await writeFile(ids, "plantsID.json");
    const plants = await getAllPlants(ids);
    await writeFile(plants, "plants.json");

    // console.timeEnd("performance");
    return Promise.resolve(plants);
  } catch (error) {
    console.log("TCL: fetchAllDataToFile -> error", error);
    return Promise.reject(error);
  } finally {
    console.timeEnd("performance");
  }

  // getAllPlantID()
  //   .then(ids => writeFile(ids, "plantsID.json"))
  //   .then(ids => getAllPlants(ids))
  //   .then(plants => writeFile(plants, "plants.json"))
  //   .then(() => console.timeEnd("performance"))
  //   .catch(() => console.timeEnd("performance"));
}

async function fetchPlantListFromDb() {
  try {
    // const ids = await db.getAllPlantIDDocs();
    const ids = idList;
    console.log("TCL: fetchPlantListFromDb -> ids", ids);

    const plants = await getAllPlants(ids);
    return Promise.resolve(plants);
  } catch (error) {
    console.log("TCL: fetchPlantListFromDb -> error", error);
    return Promise.reject(error);
  } finally {
    console.timeEnd("performance");
  }
}

async function fetchAllData() {
  try {
    const ids = await getAllPlantID();
    const plants = await getAllPlants(ids);
    console.timeEnd("performance");
    return Promise.resolve(plants);
  } catch (error) {
    console.log("TCL: fetchAllData -> error", error);
    return Promise.reject(error);
  } finally {
    console.timeEnd("performance");
  }

  // getAllPlantID()
  //   .then(ids => getAllPlants(ids))
  //   .then(() => console.timeEnd("performance"))
  //   .catch(() => console.timeEnd("performance"));
}

console.time("performance");
switch (args) {
  case "ids":
    fetchPlantIdList();
    break;

  case "plants":
    fetchPlantListFromFile();
    break;
  case "plantsdb":
    fetchPlantListFromDb();
    break;
  case "file":
    fetchAllDataToFile();
    break;

  default:
    fetchAllData();
    break;
}
