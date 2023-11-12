class Api::QuestionsController < ApplicationController
  def index
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE, STUDENT_ROLE]

    @questions = Question.where(test_id: params[:test_id])
    render json: @questions, include: :answers
  end

  def show
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE, STUDENT_ROLE]

    @question = Question.find(params[:id])
    render json: @question, include: :answers
  end

  def create
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TEACHER_ROLE]

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
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TEACHER_ROLE]

    @question = Question.find(params[:id])

    if @question.update(question_params)
      redirect_to @question.path
    else
      render json: {message: 'not updated'}, status: 400
    end
  end

  def destroy
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TEACHER_ROLE]

    @question = Question.find(params[:id])
    @question.destroy
  end

  private

  def question_params
    params.require(:question).permit(:body, :is_multiple_answer)
  end
end