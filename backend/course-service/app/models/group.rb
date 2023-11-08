class Group < ApplicationRecord
  belongs_to :course

  def path
    "/api/groups/#{id}"
  end
end
