class Job < ApplicationRecord
  has_many :homeworks
  has_many :tests

  def path
    "/api/jobs/#{id}"
  end
end
