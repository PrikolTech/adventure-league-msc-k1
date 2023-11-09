class Api::JobsController < ApplicationController
  def index
    @jobs = Job.all
    render json: @jobs
  end

  def show
    @job = Job.find(params[:id])

    render json: @job, include: :tests
  end

  def create
    @job = Job.create(job_params)
    
    if @job
      redirect_to api_job_url @job
    else
      render json: {message: 'not created'}, status: 400
    end
  end

  def update 
    @job = Job.find(params[:id])

    if @job.update(job_params)
      redirect_to api_job_url @job
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
