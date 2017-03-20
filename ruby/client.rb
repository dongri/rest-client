require 'httpclient'
require 'json'

class Client

  attr_writer :client
  
  def initialize(endpoint, header)
    @endpoint = endpoint
    @extheader = header
  end

  def get(path, query)
    self.do("GET", path, query)
  end
  
  def post(path, query)
    self.do("POST", path, query)
  end
  
  def put(path, query)
    self.do("PUT", path, query)
  end

  def delete(path, query)
    self.do("DELETE", path, query)
  end
  
  def patch(path, query)
    self.do("PATCH", path, query)
  end

  def do(method, path, query)
    uri = "#{@endpoint}#{path}"
    response = self.client.request(method, uri, query, @extheader)
    return response
  end

  def client
    @client ||= HTTPClient.new
  end
end
