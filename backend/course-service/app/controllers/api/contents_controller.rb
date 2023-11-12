class Api::ContentsController < ApplicationController
  EXTERNAL_ERROR = 'file service error'
  
  def index
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE, STUDENT_ROLE]

    @contents = Content.where(lecture_id: params[:lecture_id])
    render json: @contents
  end
  
  def show
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE, STUDENT_ROLE]

    @content = Content.find(params[:id])
    render json: @content, except: :content_type_id, include: :content_type
  end
  
  def create
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE]

    lecture = Lecture.find(params[:lecture_id])

    return render json: { message: "no file", status: 400 } unless params[:file]

    begin
      @content = Content.send_and_create(lecture, params[:file])  
    rescue => exception
      return render json: {message: EXTERNAL_ERROR, status: 500}
    end
    

    if @content
      render json: @content, except: :content_type_id, include: :content_type
    else
      unless course_type
        render json: {message: "bad request", status: 400}
      end
    end
  end

  def destroy
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE]

    @content = Content.find(params[:id])
    @content.destroy
  end
end
