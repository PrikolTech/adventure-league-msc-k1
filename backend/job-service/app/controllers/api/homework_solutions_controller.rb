class Api::HomeworkSolutionsController < ApplicationController
  EXTERNAL_ERROR = 'file service error'.freeze
  
  def index
    @solutions = HomeworkSolution.where(homework_id: params[:homework_id])
    if params.include? :user_id
      @solutions = @solutions.where(user_id: params[:user_id])
    end
    
    render json: @solutions, include: :homework_result
  end
  
  def show
    @solution = HomeworkSolution.find(params[:id])

    render json: @solution, include: :homework_result
  end

  def create
    return render json: {message: 'no user_id', status: 400} unless params.include? :user_id

    @solution = HomeworkSolution.new(user_id: params[:user_id])
    
    if @solution
      begin
        # url = HomeworkSolution.send_file(params[:file])  
        url = Intercommunication::FileService.send_file(params[:file])
        raise StandardError.new if url.nil? 
      rescue => exception
        return render json: {message: EXTERNAL_ERROR, status: 500}
      end

      @solution.file_url = url
      @solution.homework_id = params[:homework_id]
      @solution.save
      
      render json: @solution, include: :homework_result
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

    render json: @solution, include: :homework_result
  end

  private
end
