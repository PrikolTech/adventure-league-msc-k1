class Test < ApplicationRecord
  belongs_to :job
  has_many :questions

  def path
    "#{job.path}/tests/#{id}"
  end
end
