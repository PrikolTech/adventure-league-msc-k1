class SolutionAnswer < ApplicationRecord
  belongs_to :test_solution
  belongs_to :answer

  def question
    anwer.question
  end
end
