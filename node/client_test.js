const assert = require('assert');
const Client = require('./client.js');

let endpoint = "https://yourserver.herokuapp.com";
let headers = {
  "Content-Type": "application/x-www-form-urlencoded"
};

const client = new Client(endpoint, headers);
client.get("/users", {}, (error, response, body) => {
  assert.equal(null, error, "error is not null");
  assert.equal(200, response.statusCode, "statusCode error");
});
