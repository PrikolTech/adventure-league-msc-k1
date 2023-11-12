class ApplicationController < ActionController::API
  NO_PERMISSION_ERROR = 'you have no permission'

  ADMIN_ROLE = 'admin'
  STUDENT_ROLE = 'student'
  TEACHER_ROLE = 'teacher'
  TUTOR_ROLE = 'tutor'
  EMPLOYEE_ROLE = 'employee'
  ENROLLEE_ROLE = 'enrollee'

  def has_permission?(allowed_roles = [])
    return false unless params.include? :roles
    return true if params[:roles].include? ADMIN_ROLE

    params[:roles].each do |role| 
      return true if allowed_roles.include? role
    end

    false
  end

  def request_allowed?
    return false unless params.include? :user_id
    return true if has_permission?

    params[:sender_id] == params[:user_id]
  end
end
