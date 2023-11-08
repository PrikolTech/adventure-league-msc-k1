class Api::CoursesController < ApplicationController
  def index
    @courses = Course.all
    render json: @courses, except: [:education_form_id, :period_id, :course_type_id], include: {
      period: {except: :id},
      course_type: {except: :id},
      education_form: {except: :id},
    }
  end

  def show
    @course = Course.find(params['id'])
    render json: @course, except: [:education_form_id, :period_id, :course_type_id], include: {
      period: {except: :id},
      course_type: {except: :id},
      education_form: {except: :id},
    }
  end

  def create
    period = Period.create(
      starts_at: params['period']['starts_at'],
      ends_at: params['period']['ends_at']
    )

    # course_type = CourseType.where(name: params['course_type']).first
    course_type = CourseType.find_or_create_by!(name: params['course_type'])
    # render json: {message: "wrong course_type"}, status: 400 unless course_type
    
    price = params['price'].to_f if params.include? 'price'

    education_form = EducationForm.find_by(name: params['education_form'])
    return render json: {message: "wrong course_type", status: 400} unless course_type

    if @course = Course.create(
        name: params['name'],
        description: params['description'],
        price: price,
        course_type_id: course_type.id,
        period_id: period.id,
        education_form_id: education_form.id
      )
      redirect_to @course.path
    else
      unless course_type
        render json: {message: "bad request", status: 400}
      end
    end
  end

  # def update
  #   @course = Course.find(params[:id])
    
  #   if @course.update(
  #       name: params['name'],
  #       description: params['description'],
  #       price: price,
  #       course_type_id: course_type.id,
  #       period_id: period.id,
  #       education_form_id: education_form.id
  #     )
  #     redirect_to @lecture.path
  #   else
  #     render json: {message: 'not updated', status: 400}
  #   end
  # end

  def destroy
    @course = Course.find(params['id'])
    redirect_to @course.path
  end
end
