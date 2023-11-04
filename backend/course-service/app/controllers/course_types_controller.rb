class CoursesController < ApplicationController
  def index
    CourseType.create(name: 'Hello')
    @courses = CourseType.all
    render json: @courses
  end
end
