class Api::QuestionsController < ApplicationController
  def index
    @questions = Question.all
    render json: @questions 
  end

  def show
    @question = Question.find(params[:id])
    render json: @question, include: :answers
  end

  def create
    @question = Question.new(question_params)
    
    if @question
      @question.test_id = params[:test_id]
      @question.save
      redirect_to @question.path
    else
      render json: {message: 'not created'}, status: 400
    end
  end

  def update
    @question = Question.find(params[:id])

    if @question.update(question_params)
      redirect_to @question.path
    else
      render json: {message: 'not updated'}, status: 400
    end
  end

  def destroy
    @question = Question.find(params[:id])
    @question.destroy
  end

  private

  def question_params
    params.require(:question).permit(:body, :is_multiple_answer)
  end
end