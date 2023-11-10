class Homework < ApplicationRecord
  belongs_to :job
  has_many :homework_solutions

  def path
    "#{job.path}/homeworks/#{id}"
  end

  def self.send_file(file)
    url = Intercommunication::FileService.send_file(file)
    
    raise StandardError if url.nil?
    
    return url
  end
end
