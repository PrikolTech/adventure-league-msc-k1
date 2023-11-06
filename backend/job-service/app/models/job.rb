class Job < ApplicationRecord
  has_many :homeworks
  has_many :tests
end
