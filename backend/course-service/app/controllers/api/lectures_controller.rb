class Api::LecturesController < ApplicationController
  def index
    @lectures = Lecture.where(course_id: params['course_id'])
    if params[:user_id]
      courses = Course.find_by_user_id(params[user_id])
      
      @lectures = @lectures.where(course: courses, is_hidden: false)
    end
    
    render json: @lectures
  end

  def show
    @lecture = Lecture.find(params['id'])
    return { status: 403 } unless @lecture.available_for?(params[:user_id])
    
    render json: @lecture, include: [contents: {except: [:content_type_id], include: [:content_type]}]
  end

  def create
    @lecture = Lecture.create(lecture_params)

    if @lecture
      course = Course.find(params['course_id'])
      @lecture.course = course
      @lecture.save

      params['contents']&.each do |content|
        type = ContentType.find_or_create_by!(name: content['type'])
        Content.create(
          lecture_id: @lecture.id,
          url: content['url'],
          content_type_id: type.id
        )
      end

      redirect_to @lecture.path
    else
      render json: {message: "bad request", status: 400}
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
