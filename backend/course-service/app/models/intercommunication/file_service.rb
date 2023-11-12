module Intercommunication
  class FileService < BaseRetriever
    class << self
      ENDPOINT = '/api/files'
      API = ENV['FILE_SERVICE_API']

      def path
        "#{API}#{ENDPOINT}"
      end

      def send_file(file)
        # result = get(path)
        form_data = [['file', file], ['name', file.original_filename]]
        response = post(path, form_data, 'multipart/form-data')
        response_data = JSON.parse response
        puts "RESPONSE IS #{response_data}"
        return response_data['url']
      end

      # TODO: File destruction
      def destroy_file(url); end
    end
  end
end