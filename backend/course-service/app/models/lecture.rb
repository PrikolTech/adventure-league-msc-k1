class Lecture < ApplicationRecord
  belongs_to :course
  has_many :contents

  def path
    "#{course.id}/lectures/#{id}"
  end
end
