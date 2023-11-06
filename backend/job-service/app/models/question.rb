class Question < ApplicationRecord
  belongs_to :test
  has_many :answers

  def path
    "#{test.path}/questions/#{id}"
  end
end
