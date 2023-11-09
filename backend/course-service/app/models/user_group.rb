class UserGroup < ApplicationRecord
  belongs_to :group

  def self.groups_by_user(user_id)
    groups = []
    self.where(user_id: user_id).each do |u_g|
      groups << u_g.group
    end

    groups
  end
end
