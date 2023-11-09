class Answer < ApplicationRecord
  belongs_to :question

  def path
    "#{question.path}/answers/#{id}"
  end
end
