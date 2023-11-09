class Api::HomeworksController < ApplicationController
  def index
    @homeworks = Homework.where(job_id: params[:job_id])
    render json: @homeworks
  end
  
  def show
    @homework = Homework.find(params[:id])

    render json: @homework, include: :homework_solutions
  end

  def create
    @homework = Homework.new(homework_params)
    
    if @homework
      @homework.job_id = params[:job_id]
      @homework.save

      redirect_to @homework.path
    else
      render json: {message: 'not created', status: 400}
    end
  end

  def update
    @homework = Homework.find(params[:id])

    if @homework.update(homework_params)
      redirect_to @homework.path
    else
      render json: {message: 'not updated', status: 400}
    end
  end

  def delete
    @homework = Homework.find(params[:id])
    @homework.delete
  end

  private

  def homework_params
    params.require(:homework).permit(:name, :description, :text)
  end
end
