class CoursesController < ApplicationController
  def index
    CourseType.create(name: 'Hello')
    CourseType.save
    @courses = CourseType.all
    render json: @courses
  end

  def show
  end
end
