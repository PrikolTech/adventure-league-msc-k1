class Lecture < ApplicationRecord
  belongs_to :course
  has_many :contents

  def path
    "#{course.id}/lectures/#{id}"
  end

  def available_for?(user_id)
    false if user_id.nil?

    groups = course.groups
    !UserGroup.where(user_id: user_id, group: groups).empty?
  end
end
