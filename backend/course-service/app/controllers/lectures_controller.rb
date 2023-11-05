class LecturesController < ApplicationController
  def index
    course = Course.find_by(id: params['course_id'])
    return render json: {message: 'course not found'}, status: 404 if course.nil?
  
    @lectures = course.lectures
    render json: @lectures
  end

  def show
    @lecture = Lecture.find_by(id: params['id'])

    return render json: {message: 'course or lecture not found'}, status: 404 if @lecture.nil?
    
    render json: @lecture, include: [contents: {except: [:content_type_id], include: [:content_type]}]
  end

  def create
    course = Course.find_by(id: params['course_id'])
    return render json: {message: 'course not found'}, status: 404 if course.nil?

    if @lecture = Lecture.create(
        name: params['name'],
        description: params['description'],
        starts_at: params['starts_at'],
        course_id: course.id
      )
      
      params['contents']&.each do |content|
        type = ContentType.find_or_create_by!(name: content['type'])
        Content.create(
          lecture_id: @lecture.id,
          url: content['url'],
          content_type_id: type.id
        )
      end

      render json: @lecture
    else
      unless course_type
        render json: {message: "bad request"}, status: 400
      end
    end
  end
end