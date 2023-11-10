module Intercommunication
  require 'net/http'

  class BaseRetriever
    class << self
      def get(url)
        uri = URI.parse(url) 
        request = Net::HTTP::Get.new(uri)
        response = Net::HTTP.start(uri.host, uri.port) { |http|
          http.request request
        }

        response.body
      end

      def post(url, form_data, content_type='multipart/form-data')
        uri = URI.parse(url) 
        request = Net::HTTP::Post.new(uri)

        request.set_form form_data, content_type
        response = Net::HTTP.start(uri.host, uri.port) { |http|
          http.request request
        }

        response.body
      end
    end
  end
end
