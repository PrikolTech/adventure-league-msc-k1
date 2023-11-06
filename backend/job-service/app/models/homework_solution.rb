class HomeworkSolution < ApplicationRecord
  belongs_to :homework
  has_one :homework_result

  def path
    "#{homework.path}/homework_solutions/#{id}"
  end
end
