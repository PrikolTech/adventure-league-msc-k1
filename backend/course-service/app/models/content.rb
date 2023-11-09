class Content < ApplicationRecord
  belongs_to :content_type
  belongs_to :lecture

  def self.send_and_create(lecture, file)
    # TODO: Send file logic
    type = ContentType.find_or_create_by!(name: 'TEST FILE')
    
    Content.create(lecture_id: lecture.id, url: 'test/url.txt', content_type_id: type.id)
  end
end
