class Api::LecturesController < ApplicationController
  def index
    @lectures = Lecture.where(course_id: params[:course_id])
    if params[:user_id]
      courses = Course.find_by_user_id(params[:user_id])

      @lectures = @lectures.where(course: courses, is_hidden: false)
    end
    
    render json: @lectures
  end

  def show
    @lecture = Lecture.find(params[:id])
    if params[:user_id]
      user_id = params[:user_id]
      UserViews.create(user_id: user_id, lecture_id: params[:id]) if UserViews.find_by(user_id: user_id, lecture_id: params[:id]).nil?

      return { status: 403 } unless @lecture.available_for?(user_id)
    end
    
    render json: @lecture, include: [contents: {except: [:content_type_id], include: [:content_type]}]
  end

  def create
    @lecture = Lecture.create(lecture_params)

    if @lecture
      course = Course.find(params[:course_id])
      @lecture.course = course
      @lecture.save

      params['contents']&.each do |content|
        Content.send_and_create(@lecture, content[:file])
      end

      redirect_to @lecture.path
    else
      render json: {message: 'bad request', status: 400}
    end
  end

  def update
    @lecture = Lecture.find(params[:id])

    if @lecture.update(homework_params)
      redirect_to @lecture.path
    else
      render json: {message: 'not updated', status: 400}
    end
  end

  def destroy
    @lecture = Lecture.find(params['id'])
    @lecture.destroy
  end

  private

  def lecture_params
    params.require(:lecture).permit(:name, :description, :starts_at) 
  end
end
