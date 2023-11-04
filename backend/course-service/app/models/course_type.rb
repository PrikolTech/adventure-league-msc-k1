class CourseType < ApplicationRecord
  # has_many :courses

  # after_create :add_default_checklist

  # DEFAULT_TYPES = [
  #   'Базовый',
  #   'Продвинутый',
  #   'Сложный'
  # ]

  # def add_default_checklist
  #   DEFAULT_TYPES.each do |type| # Define way to get default values
  #     c_t = CourseType.create(name: type)
  #     c_t.save()
  #   end
  # end
end
