class Api::JobsController < ApplicationController
  def index
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE, STUDENT_ROLE]

    @jobs = Job.all

    # TODO: Check user access to job's lection!!!!!!!!!!!!!!!
    if params[:lecture_id]
      @jobs = Job.where(lecture_id: params[:lecture_id])
    end

    render json: @jobs, include: [:tests, :homeworks]
  end

  def show
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE, STUDENT_ROLE]

    @job = Job.find(params[:id])

    render json: @job, include: [:tests, :homeworks]
  end

  def create
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TEACHER_ROLE]

    @job = Job.create(job_params)
    
    if @job
      render json: @job, include: :tests
    else
      render json: {message: 'not created'}, status: 400
    end
  end

  def update
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TEACHER_ROLE]
    
    @job = Job.find(params[:id])

    if @job.update(job_params)
      redirect_to @job.path
    else
      render json: {message: 'not updated'}, status: 400
    end
  end

  def destroy
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TEACHER_ROLE]
    
    @job = Job.find(params[:id])
    @job.destroy
  end

  private 

  def job_params
    params.require(:job).permit(:lecture_id, :name, :deadline)
  end
end
