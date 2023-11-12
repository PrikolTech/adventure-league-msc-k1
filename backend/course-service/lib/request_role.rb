class RequestRole
  require 'net/http'

  def initialize(app)
    @app = app
  end

  def call(env)
    # Middleware method of verifying request

    token = env.fetch('HTTP_AUTHORIZATION', nil)
    sender_id, roles = verify_token(token)

    new_query = "" << env['QUERY_STRING']

    puts "SENDER = #{sender_id} | ROLES = #{roles}"

    if sender_id
      new_query << "&sender_id=#{sender_id}"
      roles.each { |role| new_query << "&roles[]=#{role}" } if roles
    end

    env['QUERY_STRING'] = new_query

    result = @app.call(env)
    result
  end

  private

  def verify_token(token)
    return [nil, nil] if token.nil?
    
    # puts request.header

    url = "#{ENV['AUTH_SERVICE_API']}/verify?access_token=#{token}"

    uri = URI.parse(url) 
    request = Net::HTTP::Get.new(uri)

    # request.set_form form_data, content_type
    response = Net::HTTP.start(uri.host, uri.port) { |http|
      http.request request
    }

    puts "RESPONSE BODY #{response.body}"
    data = JSON.parse response.body

    [data['id'], data['roles']]
  end
end
