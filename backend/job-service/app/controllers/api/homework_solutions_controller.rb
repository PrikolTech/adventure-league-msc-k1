class Api::HomeworkSolutionsController < ApplicationController
  EXTERNAL_ERROR = 'file service error'.freeze
  
  def index
    @solutions = HomeworkSolution.where(homework_id: params[:homework_id])
    render json: @solutions, include: :homework_result
  end
  
  def show
    @solution = HomeworkSolution.find(params[:id])

    render json: @solution, include: :homework_result
  end

  def create
    @solution = HomeworkSolution.new(homework_solution_params)
    
    if @solution
      begin
        url = HomeworkSolution.send_file(params[:file])  
      rescue => exception
        return render json: {message: EXTERNAL_ERROR, status: 500}
      end

      @solution.url = url
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
    @solution = HomeworkSolution.find(params[:id])
    @solution.delete
  end

  def result
    @solution = HomeworkSolution.find(params[:homework_solution_id])
    HomeworkResult.create(score: params[:score], homework_solution_id: @solution.id)

    redirect_to @solution.path
  end

  private

  def homework_solution_params
    params.require(:homework_solution).permit(:user_id)
  end
end
