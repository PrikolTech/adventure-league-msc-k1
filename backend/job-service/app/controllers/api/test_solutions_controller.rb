class Api::TestSolutionsController < ApplicationController
  
  # Формируется в момент отправки всех ответов на сервер
  def index
    @solutions = TestSolution.where(test_id: params[:test_id])
    @solutions.where(user_id: params[:user_id]) if params[:user_id]
    render json: @solutions
  end

  def show
    @solution = TestSolution.find(params[:id])
    puts "RESULT IS #{@solution.check_correctness}"

    render json: @solution, include: :solution_answers, include: :test_result
  end

  def create
    @solution = TestSolution.new(test_solution_params)

    if @solution
      @solution.test_id = params[:test_id]
      @solution.save
      params[:answers].each do |answer|
        create_answer answer
      end

      @solution.test_result.create(@solution.check_correctness)
      @solution.save

      redirect_to @solution.path
    else
      render json: {message: 'not created', status: 400} 
    end
  end

  def destroy
    @solution = TestSolution.find(params[:id])
    @solution.solution_answers.each do |answer|
      answer.destroy
    end

    @solution.destroy
  end

  private

  def create_answer(answer)
    permitted_answer = answer.permit(:answer_id)
    @solution.solution_answers.create permitted_answer
  end

  def test_solution_params
    params.require(:test_solution).permit(:user_id)
  end
end