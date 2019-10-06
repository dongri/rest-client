require 'uri'
require 'net/http'
require 'json'

class Client

  def initialize(endpoint, headers)
    uri = URI.parse(endpoint)
    @http = Net::HTTP.new(uri.host, uri.port)
    @headers = headers
  end

  def get(path, params)
    self.do(:get, path, params)
  end
  
  def post(path, params)
    self.do(:post, path, params)
  end
  
  def put(path, params)
    self.do(:put, path, params)
  end

  def delete(path, params)
    self.do(:delete, path, params)
  end
  
  def patch(path, params)
    self.do(:patch, path, params)
  end

  def do(method, path, params)
    case method
    when :get then
      full_path = path_with_params(path, params)
      request = Net::HTTP::Get.new(full_path)
    when :post then
      request = Net::HTTP::Post.new(path)
      # request.set_form_data(params)
    when :put then
      request = Net::HTTP::Put.new(path)
      # request.set_form_data(params)
    when :delete then
      full_path = path_with_params(path, params)
      request = Net::HTTP::Get.new(full_path)
    when :patch then
      request = Net::HTTP::Patch.new(path)
      # request.set_form_data(params)
    end
    for header in @headers do
      request.add_field(header[0], header[1])
      # if header[0] == "Content-Type" && header[1] == "application/json"
        # request.body = params.to_json
      # end 
    end
    request.set_form_data(params)
    # request.body = params.to_json
    response = @http.request(request)
    return response
  end

  def path_with_params(path, params)
    encoded_params = URI.encode_www_form(params)
    [path, encoded_params].join("?")
  end

end
