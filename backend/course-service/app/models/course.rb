class Course < ApplicationRecord
  belongs_to :course_type
  belongs_to :period
  belongs_to :education_form

  has_many :lectures
end
