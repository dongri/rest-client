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

client.post("/users", {"name": "dongri", "email": "dongri@origami.com"}, (error, response, body) => {
  assert.equal(null, error, "error is not null");
  assert.equal(201, response.statusCode, "statusCode error");
  assert.equal("dongri", JSON.parse(body).name);
});

client.put("/users/1", {"name": "dongri2", "email": "dongri2@origami.com"}, (error, response, body) => {
  assert.equal(null, error, "error is not null");
  assert.equal(200, response.statusCode, "statusCode error");
  assert.equal("dongri2", JSON.parse(body).name);
});

client.delete("/users/1", {}, (error, response, body) => {
  assert.equal(null, error, "error is not null");
  assert.equal(200, response.statusCode, "statusCode error");
});
