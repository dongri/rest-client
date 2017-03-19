require 'httpclient'
require 'json'

class Client

  attr_writer :client
  
  def initialize(endpoint, header)
    @endpoint = endpoint
    @extheader = header
  end

  def get(path, query)
    self.do({
      :method => "GET",
      :path => path,
      :query => query
    })
  end
  
  def post(path, query)
    self.do({
      :method => "POST",
      :path => path,
      :query => query
    })
  end
  
  def put(path, query)
    self.do({
      :method => "PUT",
      :path => path,
      :query => query
    })
  end

  def delete(path, query)
    self.do({
      :method => "DELETE",
      :path => path,
      :query => query
    })
  end
  
  def patch(path, query)
    self.do({
      :method => "PATCH",
      :path => path,
      :query => query
    })
  end

  def get_method(method)
    return method
  end

  def get_uri(path)
    "#{@endpoint}#{path}"
  end

  def get_query(query)
    return query
  end

  def do(attrs)
    client = self.client
    method = self.get_method(attrs[:method])
    uri = self.get_uri(attrs[:path])
    query = self.get_query(attrs[:query])
    response = client.request(method, uri, query, @extheader)
    return response
  end

  def client
    @client ||= HTTPClient.new
  end
end
