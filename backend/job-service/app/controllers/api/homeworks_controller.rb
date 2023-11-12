class Api::HomeworksController < ApplicationController
  def index
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE, STUDENT_ROLE]

    # TODO: check access to lection
    @homeworks = Homework.where(job_id: params[:job_id])
    render json: @homeworks
  end
  
  def show
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TEACHER_ROLE]

    @homework = Homework.find(params[:id])

    render json: @homework, include: :homework_solutions
  end

  def create
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TEACHER_ROLE]

    @homework = Homework.new(homework_params)
    
    if @homework
      @homework.job_id = params[:job_id]
      @homework.save

      render json: @homework, include: :homework_solutions
    else
      render json: {message: 'not created', status: 400}
    end
  end

  def update
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TEACHER_ROLE]

    @homework = Homework.find(params[:id])

    if @homework.update(homework_params)
      render json: @homework, include: :homework_solutions
    else
      render json: {message: 'not updated', status: 400}
    end
  end

  def destroy
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TEACHER_ROLE]

    @homework = Homework.find(params[:id])
    @homework.destroy
  end

  private

  def homework_params
    params.require(:homework).permit(:name, :description, :text)
  end
end
