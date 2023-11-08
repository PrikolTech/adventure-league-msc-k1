class Group < ApplicationRecord
  belongs_to :course
  has_many :user_group

  MAX_USERS = 20

  def path
    "/api/groups/#{id}"
  end

  def add_user(user_id)
    return nil if full?
    
    UserGroup.create(user_id: user_id, group_id: id)
  end

  def full?
    user_group.count >= Group::MAX_USERS
  end
end
