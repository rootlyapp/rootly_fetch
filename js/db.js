const loki = require("lokijs");

let idsCollection, plantsCollection;
const db = new loki("plants.json", {
  // autoload: true,
  autosave: true
  // autoloadCallback: database => {
  //   idsCollection =
  //     database.getCollection("ids") || database.addCollection("ids");
  //   plantsCollection =
  //     database.getCollection("plants") || database.addCollection("plants");
  // }
});
// const idsCollection = db.getCollection("ids") || db.addCollection("ids");
// const plantsCollection =
//   db.getCollection("plants") || db.addCollection("plants");

function loadDatabase() {
  return new Promise((resolve, reject) => {
    db.loadDatabase({}, err => {
      if (err) {
        console.log("TCL: loadDatabase -> err", err);
        reject(err);
      } else {
        idsCollection = db.getCollection("ids") || db.addCollection("ids");
        plantsCollection =
          db.getCollection("plants") || db.addCollection("plants");

        resolve();
      }
    });
  });
}

async function addPlantToCollection(plant) {
  try {
    // const plant = plantsCollection.find({ id: plant.id });
    const _plant = plantsCollection.find({ id: plant.id })[0];
    if (!_plant) {
      plantsCollection.insert(plant);
      console.log(
        "Adding new plant to DB: [%i] %s",
        plant.id,
        plant.scientific_name
      );
    }
    else {
      plantsCollection.update({ ..._plant, ...plant });
      console.log(
        "Updating plant to DB: [%i] %s",
        plant.id,
        plant.scientific_name
      );
    }
    // await loadDatabase();
    // plantsCollection.find({ id: plant.id })
    //   ? plantsCollection.insert(plant)
    //   : plantsCollection.update(plant);
    // console.log(
    //   "Adding new plant to DB: [%i] %s",
    //   plant.id,
    //   plant.scientific_name
    // );
    return Promise.resolve(plant.id);
  } catch (error) {
    console.log("TCL: addPlantToCollection -> error", error);
    return Promise.reject(error);
  }
}

async function getAllPlantIDDocs() {
  try {
    await loadDatabase();
    const response = idsCollection.get(1).list;
    return Promise.resolve(response);
  } catch (error) {
    console.log(__dirname);
    console.log("TCL: getAllPlantIDDocs -> error", error);
    return Promise.reject(error);
  }
}

async function saveIdCollection(idList) {
  try {
    await loadDatabase();

    console.log("\nSAVING ID LIST TO DB\n");
    let doc = idsCollection.get(1);
    if (!doc) {
      console.log("Creating ID document\n\n");

      doc = { name: "id", list: idList };
      idsCollection.insert(doc);
    } else {
      console.log("Updating ID document\n\n");

      doc.list = [...new Set(doc.list.concat(idList))];
      idsCollection.update(doc);
    }

    return Promise.resolve(true);
  } catch (error) {
    console.log("TCL: saveIdCollection -> error", error);
    return Promise.reject(error);
  }
  // loadDatabase()
  //   .then(() => {
  //     console.log("\nSAVING ID LIST TO DB\n");

  //     let doc = idsCollection.get(1);

  //     if (!doc) {
  //       console.log("Creating ID document\n\n");

  //       doc = { name: "id", list: idList };
  //       idsCollection.insert(doc);
  //     } else {
  //       console.log("Updating ID document\n\n");

  //       doc.list = [...new Set(doc.list.concat(idList))];
  //       idsCollection.update(doc);
  //     }
  //   })
  //   .catch(err => console.err(err));

  // let doc = idsCollection.get(1) || { name: "id", list: [] };
  // doc.list;
  // ids = [...new Set(ids.concat(idList))];
}

module.exports = {
  addPlantToCollection: addPlantToCollection,
  saveIdCollection: saveIdCollection,
  getAllPlantIDDocs: getAllPlantIDDocs,
  loadDatabase: loadDatabase
};
