class Api::LecturesController < ApplicationController
  def index
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE, STUDENT_ROLE]

    @lectures = Lecture.where(course_id: params[:course_id])
    if request_allowed?
      user_id = params[:user_id]

      courses = Course.find_by_user_id(user_id)

      @lectures = @lectures.where(course: courses, is_hidden: false)
      
      lectures_with_views = []
      @lectures.each {|lecture| lectures_with_views << lecture.attributes.merge({is_viewed: !UserViews.where(user_id: user_id, lecture_id: lecture.id).empty?})}
      @lectures = lectures_with_views
    end

    render json: @lectures
  end

  def show
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE, STUDENT_ROLE]

    @lecture = Lecture.find(params[:id])
    if request_allowed?
      user_id = params[:user_id]

      return { status: 403 } unless @lecture.available_for?(user_id)
      UserViews.create(user_id: user_id, lecture_id: params[:id]) if UserViews.find_by(user_id: user_id, lecture_id: params[:id]).nil?
    end
    
    render json: @lecture, include: [contents: {except: [:content_type_id], include: [:content_type]}]
  end

  def create
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE]

    @lecture = Lecture.create(lecture_params)

    if @lecture
      course = Course.find(params[:course_id])
      @lecture.course = course
      @lecture.save

      params['contents']&.each do |content|
        Content.send_and_create(@lecture, content[:file])
      end

      render json: @lecture, include: [contents: {except: [:content_type_id], include: [:content_type]}]
    else
      render json: {message: 'bad request', status: 400}
    end
  end

  def update
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE]

    @lecture = Lecture.find(params[:id])

    if @lecture.update(lecture_params)
      render json: @lecture, include: [contents: {except: [:content_type_id], include: [:content_type]}]
    else
      render json: {message: 'not updated', status: 400}
    end
  end

  def destroy
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE]

    @lecture = Lecture.find(params['id'])
    @lecture.destroy
  end

  private

  def lecture_params
    params.require(:lecture).permit(:name, :description, :starts_at) 
  end
end
