class CourseTypesController < ApplicationController
  def index
    @courses = CourseType.all
    render json: @courses
  end
end
