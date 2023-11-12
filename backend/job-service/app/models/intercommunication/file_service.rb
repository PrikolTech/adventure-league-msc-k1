module Intercommunication
  class FileService < BaseRetriever
    class << self
      ENDPOINT = '/api/files'
      HOST = '172.20.0.1'
      PORT = 3008

      def path
        "http://#{HOST}:#{PORT}#{ENDPOINT}"
      end

      def send_file(file)
        # result = get(path)
        form_data = [['file', file], ['name', file.original_filename]]
        response = post(path, form_data, 'multipart/form-data')
        response_data = JSON.parse response
        puts "RESPONSE #{response_data}"
        
        return response_data['url']
      end

      # TODO: File destruction
      def destroy_file(url); end
    end
  end
end
