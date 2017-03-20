request = require('request');

class Client {

  constructor(endpoint, headers) {
    this.endpoint = endpoint;
    this.headers = headers;
  }
  
  get(path, query, callback) {
    this.do('GET', path, query, callback);
  }

  post(path, query, callback) {
    this.do('POST', path, query, callback);
  }

  put(path, query, callback) {
    this.do('PUT', path, query, callback);
  }

  delete(path, query, callback) {
    this.do('DELETE', path, query, callback);
  }

  patch(path, query, callback) {
    this.do('PATCH', path, query, callback);
  }

  do(method, path, params, callback) {
    let options = {
      method: method,
      headers: this.headers,
      uri: this.endpoint + path,
      json: params
    };
    request(options, (error, response, body) => {
      callback(error, response, body)
    });
  }

}

module.exports = Client
