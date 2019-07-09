
const axios = require("axios");


const Token = "eWZMaTRwb0s3akNkMGJjMDU3Mm01UT09"

const PerPage = 2000

// TotalPages Total available pages
const TotalPages = 74

// URL Endpoint
const URL = "https://trefle.io/api/plants"

let list = [];
let counter = 0;

(async function test(){
    for (let i = 1; i <= TotalPages; i++) {
  

        const resp = await axios.get(URL, {
            params: {
              token: Token,
              page: i,
              page_size: PerPage,
            }
          });
    
          const items = resp.data.map(it => {
            //   console.log(it.complete_data);
              if(it.complete_data){
                  counter++
              }
              return it.id
          });
          
          list = list.concat(items);
          console.log(`Page ${i}\\${TotalPages} - TOTAL:${list.length}`);
          
    }
    console.log(`TOTAL:${list.length}`)
    console.log(`COUNT:${counter}`)
})()
