class Course < ApplicationRecord
  belongs_to :course_type
  belongs_to :period
  belongs_to :education_form

  has_many :lectures
  has_many :groups

  MAX_GROUPS = 3

  def path
    "/api/#{id}"
  end

  def new_group_name
    "#{prefix}-#{(groups.count + 1).to_s.rjust 2, '0'}"
  end

  def add_group
    return nil if groups.count >= MAX_GROUPS

    Group.create(name: new_group_name, course_id: id)
  end

  def started?
    period.starts_at <= DateTime.now
  end

  def self.find_by_user_id(user_id)
    courses = []

    course if user_id.nil?

    groups = UserGroup.groups_by_user(user_id)
    groups.each { |group| courses << group.course }

    courses
  end
end
