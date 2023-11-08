class Api::CoursesController < ApplicationController
  USER_ADDED_ERROR = 'user is already added'

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

    course_type = CourseType.find_or_create_by!(name: params['course_type'])
    
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

  def add_user
    course = Course.find(params[:course_id])
    groups = Group.where(course_id: params[:course_id])
    user_id = params[:user_id]

    return render json: {message: USER_ADDED_ERROR, status: 400} unless UserGroup.where(group: groups, user_id: user_id).empty?


    is_added = false
    groups.each do |group|
      user_group = group.add_user(user_id)
      return render json: {group: group, status: 201}, include: :course if user_group
    end

    if groups.count < Course::MAX_GROUPS
      @group = course.add_group
      @group.add_user(user_id)
    end

    if @group
      render json: {group: @group, status: 201}, include: :course
    else
      render json: {message: 'can not add user', status: 400}
    end
  end
end
