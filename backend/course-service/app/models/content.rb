class Content < ApplicationRecord
  belongs_to :content_type
  belongs_to :lecture
end
