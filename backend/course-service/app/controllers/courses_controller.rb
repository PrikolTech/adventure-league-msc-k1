class CoursesController < ApplicationController
  def index
    @courses = CourseType.all
  end
end
