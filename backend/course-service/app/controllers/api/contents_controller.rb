class Api::ContentsController < ApplicationController
  def index
    @contents = Content.where(lecture_id: params['lecture_id'])
    render json: {contents: @contents, status: 200}
  end
  
  def show
    @content = Content.find(params['id'])
    render json: @content
  end
  
  def create
    lecture = Lecture.find(params['lecture_id'])

    @content = Content.send_and_create(lecture, nil)

    if @content
      render json: @content, include: [:content_type]
    else
      unless course_type
        render json: {message: "bad request", status: 400}
      end
    end
  end

  private

  def content_params
    params.require(:content)
  end
end
