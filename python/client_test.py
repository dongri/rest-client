import unittest
from client import Client

class TestClient(unittest.TestCase):

  endpoint = "https://yourserver.herokuapp.com"
  # endpoint = 'http://localhost:8080'

  def test_get(self):
    headers = {
      'Content-Type': 'application/json'
    }
    client = Client(self.endpoint, headers)
    params = {
      'name': 'dongri'
    }
    body = client.get('/', params)
    print(body)

  def test_post(self):
    headers = {
      'Content-Type': 'application/x-www-form-urlencoded'
    }
    client = Client(self.endpoint, headers)
    params = {
        'name': 'dongri'
    }
    body = client.post('/', params)
    print(body)

  def test_post_json(self):
    headers = {
        'Content-Type': 'application/json'
    }
    client = Client(self.endpoint, headers)
    params = {
        'name': 'dongri'
    }
    body = client.post('/', params)
    print(body)

  def test_put(self):
    headers = {
        'Content-Type': 'application/x-www-form-urlencoded'
    }
    client = Client(self.endpoint, headers)
    params = {
        'name': 'dongri'
    }
    body = client.put('/', params)
    print(body)


  def test_delete(self):
    headers = {
        'Content-Type': 'application/x-www-form-urlencoded'
    }
    client = Client(self.endpoint, headers)
    params = {
        'name': 'dongri'
    }
    body = client.delete('/', params)
    print(body)

  def test_patch(self):
    headers = {
        'Content-Type': 'application/x-www-form-urlencoded'
    }
    client = Client(self.endpoint, headers)
    params = {
        'name': 'dongri'
    }
    body = client.patch('/', params)
    print(body)



if __name__ == '__main__':
    unittest.main()
