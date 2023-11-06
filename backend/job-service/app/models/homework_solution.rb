class HomeworkSolution < ApplicationRecord
  belongs_to :homework
  has_one :homework_result
end
