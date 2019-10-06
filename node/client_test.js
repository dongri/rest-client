const assert = require('assert');
const Client = require('./client.js');

// let endpoint = "https://yourserver.herokuapp.com";
let endpoint = "http://localhost:8080";

var headers = {
  "Content-Type": "application/x-www-form-urlencoded"
};

var client = new Client(endpoint, headers);

client.get("/", {"name": "dongri"}, (error, response, body) => {
  assert.equal(null, error, "error is not null");
  assert.equal(200, response.statusCode, "statusCode error");
  console.log(body);
});

client.post("/", {"name": "dongri"}, (error, response, body) => {
  assert.equal(null, error, "error is not null");
  assert.equal(200, response.statusCode, "statusCode error");
  console.log(body);
});

client.put("/", {"name": "dongri"}, (error, response, body) => {
  assert.equal(null, error, "error is not null");
  assert.equal(200, response.statusCode, "statusCode error");
  console.log(body);
});

client.delete("/", {"name": "dongri"}, (error, response, body) => {
  assert.equal(null, error, "error is not null");
  assert.equal(200, response.statusCode, "statusCode error");
  console.log(body);
});

headers["Content-Type"] = "application/json"
client = new Client(endpoint, headers);
client.post("/", { "name": "dongri" }, (error, response, body) => {
  assert.equal(null, error, "error is not null");
  assert.equal(200, response.statusCode, "statusCode error");
  console.log(body);
});
