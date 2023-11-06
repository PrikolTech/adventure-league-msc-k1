class Api::HomeworksController < ApplicationController
  def index
    @homeworks = Homework.all
    render json: @homeworks
  end
  
  def create

  end

  def delete
    @homework = Homework.find(params[:id])
    @homework.delete
  end

  # def update
  # end
end
