class Homework < ApplicationRecord
  belongs_to :job
  has_many :homework_solutions

  def path
    "#{job.path}/homeworks/#{id}"
  end
end
