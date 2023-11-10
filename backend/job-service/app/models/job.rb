class Job < ApplicationRecord
  has_many :homeworks
  has_many :tests
  belongs_to :job_type

  def path
    "/api/jobs/#{id}"
  end
end
