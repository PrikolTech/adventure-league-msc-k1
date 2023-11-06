class Api::TestsController < ApplicationController
  def index
    @tests = Test.where(job_id: params[:job_id])
    render json: @tests
  end

  def show
    @test = Test.find(params[:id])
    render json: @test, include: :questions
  end

  def create
    @test = Test.new(test_params)
    
    if @test
      @test.job_id = params[:job_id]
      @test.save

      redirect_to @test.path
    else
      render json: {message: 'not created'}, status: 400
    end
  end

  def update 
    @test = Test.find(params[:id])

    if @test.update(test_params)
      redirect_to @test.path
    else
      render json: {message: 'not updated'}, status: 400
    end
  end

  def destroy
    @test = Test.find(params[:id])
    @test.destroy
  end

  private

  def test_params
    params.require(:test).permit(:name)
  end
end