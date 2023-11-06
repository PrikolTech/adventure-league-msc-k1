class Api::AnswersController < ApplicationController
  CORRECT_EXIST_ERROR = 'correct answer already exist'.freeze

  def index
    @answers = Answer.where(question_id: params[:question_id])
    render json: @answers
  end

  def show
    @answer = Answer.find(params[:id])
    render json: @answer
  end

  def create
    return render json: {message: CORRECT_EXIST_ERROR, status: 409} unless create_available?

    @answer = Answer.new(answer_params)

    if @answer
      @answer.question_id = params[:question_id]
      @answer.save

      redirect_to @answer.path
    else
      render json: {message: 'not created', status: 400}
    end
  end

  def updated
    @answer = Answer.find(params[:id])

    return render json: {message: CORRECT_EXIST_ERROR, status: 409} unless create_available?

    if @answer.update(answer_params)
      redirect_to @answer.path
    else
      render json: {message: 'not updated', status: 400}
    end
  end

  def destroy
    @answer = Answer.find(params[:id])
    @answer.destroy
  end

  private

  def create_available?
    question = Question.find(params[:question_id])
    another_answers = question.answers

    if !question.is_multiple_answer && answer_params[:is_correct]
      is_correct_exists = false
      another_answers.each do |answer|
        is_correct_exists ||= answer.is_correct
      end

      return false if is_correct_exists
    end

    true
  end

  def answer_params
    params.require(:answer).permit(:body, :is_correct)
  end
end