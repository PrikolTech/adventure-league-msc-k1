class Content < ApplicationRecord
  belongs_to :content_type
  belongs_to :lecture

  def self.send_and_create(lecture, file)
    type = ContentType.find_or_create_by!(name: 'TEST FILE')
    url = Intercommunication::FileService.send_file(file)
    
    raise StandardError if url.nil?
    
    Content.create(lecture_id: lecture.id, url: url, content_type_id: type.id)
  end

  def path
    "#{lecture.path}/contents/#{id}"
  end
end
