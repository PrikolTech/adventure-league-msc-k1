class Api::HomeworkSolutionsController < ApplicationController
  def index
    @solutions = HomeworkSolution.where(homework_id: params[:homework_id])
    render json: @solutions
  end
  
  def show
    @solution = HomeworkSolution.find(params[:id])

    render json: @solutions, include: :homework_result
  end

  def create
    @solution = HomeworkSolution.new(homework_solution_params)
    
    if @solution
      @solution.homework_id = params[:homework_id]
      @solution.save

      redirect_to @solution.path
    else
      render json: {message: 'not created', status: 400}
    end
  end

  def update
    @solution = HomeworkSolution.find(params[:id])

    if @solution.update(homework_solution_params)
      redirect_to @solution.path
    else
      render json: {message: 'not updated', status: 400}
    end
  end

  def delete
    @homework = Homework.find(params[:id])
    @homework.delete
  end

  private

  def homework_solution_params
    params.require(:homework_solution).permit(:user_id, :file_url)
  end
end
