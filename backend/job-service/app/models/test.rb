class Test < ApplicationRecord
  belongs_to :job
  has_many :questions
end
