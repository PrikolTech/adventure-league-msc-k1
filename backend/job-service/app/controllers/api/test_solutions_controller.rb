class Api::TestSolutionsController < ApplicationController
  
  # Формируется в момент отправки всех ответов на сервер
  def index
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE, STUDENT_ROLE]

    @solutions = []
    if has_permission? [TUTOR_ROLE, TEACHER_ROLE]
      @solutions = TestSolution.where(test_id: params[:test_id])
      @solutions = @solutions.where(user_id: params[:user_id]) if params[:user_id]
    else
      @solutions = @solutions.where(user_id: params[:sender_id]) if params[:sender_id] 
    end

    render json: @solutions, include: [:test_result]
  end

  def show
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE, STUDENT_ROLE]

    @solution = TestSolution.find(params[:id])

    return json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE] && request_allowed?

    render json: @solution, include: [:test_result, :solution_answers]
  end

  def create
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE, STUDENT_ROLE]

    @solution = TestSolution.new(test_solution_params)

    if @solution
      @solution.test_id = params[:test_id]
      @solution.save
      params[:answers].each do |answer|
        create_answer answer
      end
      @solution.save

      # @solution.test_result_id = TestResult.create(score: @solution.check_correctness, test_solution_id = @solution.id)
      test_result = TestResult.create(score: @solution.check_correctness, test_solution_id: @solution.id)

      # test_result
      # redirect_to @solution.path
      render json: @solution, include: [:test_result, :solution_answers]
    else
      render json: {message: 'not created', status: 400} 
    end
  end

  def destroy
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? []

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