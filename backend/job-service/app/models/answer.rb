class Answer < ApplicationRecord
  belongs_to :question

  def path
    "#{question.path}/answer/#{id}"
  end
end
