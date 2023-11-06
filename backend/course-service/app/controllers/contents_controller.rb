class ContentsController < ApplicationController
  def create
    lecture = Lecture.find_by(id: params['id'])
    return render json: {message: 'lecture not found'}, status: 404 if lecture.nil?

    @contents = []

    params['contents']&.each do |content|
      type = ContentType.find_or_create_by!(name: content['type'])
      @contents << Content.create(
        lecture_id: lecture.id,
        url: content['url'],
        content_type_id: type.id
      )
    end

    if @contents

      render json: @contents, except: [:content_type_id], include: [:content_type]
    else
      unless course_type
        render json: {message: "bad request"}, status: 400
      end
    end
  end
end
