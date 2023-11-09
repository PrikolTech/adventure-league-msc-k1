class Question < ApplicationRecord
  belongs_to :test
  has_many :answers

  def path
    "#{test.path}/questions/#{id}"
  end

  def correct_answers
    result = []

    answers.each do |answer|
      result.push answer if answer.is_correct
    end

    result
  end
end
