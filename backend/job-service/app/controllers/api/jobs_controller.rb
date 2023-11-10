class Api::JobsController < ApplicationController
  def index
    @jobs = Job.all

    # Intercommunication::BaseRetriever.get('http://172.20.0.1:3003/api/courses')
    if params[:lecture_id]
      @jobs = @jobs.where(lecture_id: params[:lecture_id])
    end

    render json: @jobs, include: [:tests, :homeworks]
  end

  def show
    @job = Job.find(params[:id])

    render json: @job, include: :tests
  end

  def create
    @job = Job.create(job_params)
    
    if @job
      render json: @job, include: :tests
    else
      render json: {message: 'not created'}, status: 400
    end
  end

  def update 
    @job = Job.find(params[:id])

    if @job.update(job_params)
      redirect_to @job.path
    else
      render json: {message: 'not updated'}, status: 400
    end
  end

  def destroy
    @job = Job.find(params[:id])
    @job.destroy
  end

  private 

  def job_params
    params.require(:job).permit(:lecture_id, :name, :deadline)
  end
end
