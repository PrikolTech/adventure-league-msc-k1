class Lecture < ApplicationRecord
  belongs_to :course
  has_many :contents
end
