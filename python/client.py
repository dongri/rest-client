import urllib.request
import json

class Client:

  def __init__(self, base_url, headers):
    self.__base_url = base_url
    self.__headers = headers

  def get(self, path, params):
    return self.do('GET', path, params)

  def post(self, path, params):
    return self.do('POST', path, params)

  def put(self, path, params):
    return self.do('PUT', path, params)

  def delete(self, path, params):
    return self.do('DELETE', path, params)

  def patch(self, path, params):
    return self.do('PATCH', path, params)

  def do(self, method, path, params):
    if method == 'GET' or method == 'DELETE':
      url = '{}{}?{}'.format(self.__base_url, path, urllib.parse.urlencode(params))
      request = urllib.request.Request(
        url, headers=self.__headers, method=method)
    else:
      for k, v in self.__headers.items():
        if k == "Content-Type" and v == "application/json":
          params = json.dumps(params).encode('utf-8')
        else:
          params = urllib.parse.urlencode(params).encode('utf-8')
      url = '{}{}'.format(self.__base_url, path)
      request = urllib.request.Request(
        url, data=params, headers=self.__headers, method=method)
    with urllib.request.urlopen(request) as response:
        response_body = response.read().decode("utf-8")
        return response_body
