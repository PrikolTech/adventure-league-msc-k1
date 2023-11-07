require 'sinatra'
require 'sinatra/json'
require 'fileutils'

ALLOWED_EXTENSION = %w(
 txt
 pdf
 mp4
 doc
 docx
 pptx
 ppt
 png
 jpeg
).freeze

MAX_SIZE = 2 * 2**20 # 2 MB

STORAGE_PATH = './storage/content'

get '/' do
  # Files send html form (DEBUG ONLY)
  erb :send
end

get '/api/files' do
  # Returns list of files (DEBUG ONLY)

  files = Dir.glob("./storage/content/*/*.*").map{ |f| "#{f.split('/')[2..-1].join('/')}" }
  return json(files.to_a)
end

get '/api/files/*/*' do
  # Returns file by path and name
  file_path, filename = params['splat']
  full_file_path = "#{STORAGE_PATH}/#{file_path}/#{filename}" 

  return json status: 404, message: 'no such file' unless File.exist?(full_file_path)

  send_file full_file_path, :filename => filename, :type => 'Application/octet-stream'
end

post '/api/files' do
  # Saves file and returns a file url
  tempfile = params[:file][:tempfile]
  filename = params[:file][:filename]
  extension = filename.split('.')[-1]
  size = File.size(tempfile)
  
  
  return json :message => 'wrong file extension', :status => 400 unless ALLOWED_EXTENSION.include? extension
  return json :message => 'file size is too large', :status => 400 if size > MAX_SIZE

  date_time_pref = DateTime.now.strftime('%m%d%Y%H%M%S')
  
  path = "#{STORAGE_PATH}/#{date_time_pref}"

  Dir.mkdir path
  FileUtils.cp(tempfile.path, "#{path}/#{filename}")
  
  json :url => "#{date_time_pref}/#{filename}"
end
