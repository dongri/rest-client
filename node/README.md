# Usage

```
const Client = require('./client.js');

let endpoint = "https://yourserver.herokuapp.com";
let headers = {
  "Content-Type": "application/x-www-form-urlencoded"
};

const client = new Client(endpoint, headers);

client.post("/users", {"name": "dongri", "email": "dongri@origami.com"}, (error, response, body) => {
  console.log(error);
  console.log(response.statusCode);
  console.log(JSON.parse(body));
});
```
