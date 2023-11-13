class Api::CoursesController < ApplicationController
  USER_ADDED_ERROR = 'user is already added'
  COURSE_STARTED_ERROR = 'course is already started'

  def index
    # puts "PARAMS #{params}"
    @courses = Course.all

    if has_permission? []
      if params[:user_id]
        @courses = Course.find_by_user_id(user_id: params[:user_id])
      end
    else
      if request_allowed?
        @courses = Course.find_by_user_id(params[:user_id])
      end
    end

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
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE]

    period = Period.create(
      starts_at: params['period']['starts_at'],
      ends_at: params['period']['ends_at']
    )

    course_type = CourseType.find_or_create_by!(name: params['course_type'])
    
    price = params['price'].to_f if params.include? 'price'

    education_form = EducationForm.find_or_create_by!(name: params['education_form'])
    return render json: {message: "wrong course_type", status: 400} unless course_type

    if @course = Course.create(
        name: params['name'],
        description: params['description'],
        price: price,
        course_type_id: course_type.id,
        period_id: period.id,
        education_form_id: education_form.id,
        prefix: params[:prefix]
      )
      render json: @course, except: [:education_form_id, :period_id, :course_type_id], include: {
        period: {except: :id},
        course_type: {except: :id},
        education_form: {except: :id},
      }
    else
      unless course_type
        render json: {message: "bad request", status: 400}
      end
    end
  end

  def update
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE]

    @course = Course.find(params[:id])
    
    if @course.update(course_params)
      redirect_to @course.path
    else
      render json: {message: 'not updated', status: 400}
    end
  end

  def destroy
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE]

    @course = Course.find(params['id'])
    redirect_to @course.path
  end

  def add_user
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [EMPLOYEE_ROLE]

    course = Course.find(params[:course_id])
    groups = Group.where(course_id: params[:course_id])
    user_id = params[:user_id]
    
    # Errors Handing
    return render json: {message: COURSE_STARTED_ERROR, status: 400} if course.started?
    return render json: {message: USER_ADDED_ERROR, status: 400} unless UserGroup.where(group: groups, user_id: user_id).empty?

    # Adding user to group
    groups.each do |group|
      user_group = group.add_user(user_id)
      return render json: {group: group, status: 201}, include: :course if user_group
    end

    # Adding user to a new group
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

  private

  def course_params
    params.require(:course).permit(:name, :description, :prefix, :price)
  end
end
