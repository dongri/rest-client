request = require('request');

class Client {
  static _contentTypeJSON = 'application/json';
  static _contentTypeForm = 'application/x-www-form-urlencoded';

  static _methodGet    = 'GET';
  static _methodPost   = 'POST';
  static _methodPut    = 'PUT';
  static _methodDelete = 'DELETE';
  static _methodPatch  = 'PATH';

  constructor(endpoint, headers) {
    this.endpoint = endpoint;
    this.headers = headers;
  }
  
  get(path, query, callback) {
    this.do(Client._methodGet, path, query, callback);
  }

  post(path, query, callback) {
    this.do(Client._methodPost, path, query, callback);
  }

  put(path, query, callback) {
    this.do(Client._methodPut, path, query, callback);
  }

  delete(path, query, callback) {
    this.do(Client._methodDelete, path, query, callback);
  }

  patch(path, query, callback) {
    this.do(Client._methodPatch, path, query, callback);
  }

  do(method, path, params, callback) {
    let options = {
      method: method,
      headers: this.headers,
      uri: this.endpoint + path,
    };
    if ([Client._methodGet, Client._methodDelete].includes(method)) {
      options["qs"] = params;
    }
    if ([Client._methodPost, Client._methodPut, Client._methodPatch].includes(method)) {
      let contentType = this.headers["Content-Type"];
      if (contentType == Client._contentTypeJSON) {
        options["json"] = params;
      }
      if (contentType == Client._contentTypeForm) {
        options["form"] = params;
      }
    }
    request(options, (error, response, body) => {
      callback(error, response, body)
    });
  }

}

module.exports = Client
