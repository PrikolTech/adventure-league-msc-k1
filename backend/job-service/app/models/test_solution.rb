class TestSolution < ApplicationRecord
  belongs_to :test
  has_many :solution_answers
end
