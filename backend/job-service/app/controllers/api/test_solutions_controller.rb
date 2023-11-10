class Api::TestSolutionsController < ApplicationController
  
  # Формируется в момент отправки всех ответов на сервер
  def index
    @solutions = TestSolution.where(test_id: params[:test_id])
    @solutions = @solutions.where(user_id: params[:user_id]) if params[:user_id]
    render json: @solutions
  end

  def show
    @solution = TestSolution.find(params[:id])

    render json: @solution, include: [:test_result, :solution_answers]
  end

  def create
    @solution = TestSolution.new(test_solution_params)

    puts "AAAAAAAAAAAAAA"
    if @solution
      @solution.test_id = params[:test_id]
      @solution.save
      puts "AAAAAAAAAAAAAA"
      params[:answers].each do |answer|
        create_answer answer
      end
      @solution.save

      # @solution.test_result_id = TestResult.create(score: @solution.check_correctness, test_solution_id = @solution.id)
      test_result = TestResult.create(score: @solution.check_correctness, test_solution_id: @solution.id)

      puts "TEST RESULT #{test_result.score}"
      puts "TEST RESULT #{test_result.test_solution}"
      puts "TEST RESULT #{@solution.test_result}"
      # test_result
      # redirect_to @solution.path
      render json: @solution, include: [:test_result, :solution_answers]
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