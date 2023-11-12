class Api::GroupsController < ApplicationController
  def index
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE]

    @groups = Group.all
    render json: @groups
  end

  def show
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TUTOR_ROLE, TEACHER_ROLE, STUDENT_ROLE]

    @group = Group.find(params[:id])

    if UserGroup.where(user_id: params[:sender_id], group_id: @group_id).empty?
      return json: {status: 403, message: NO_PERMISSION_ERROR}
    end

    render json: @group, include: {user_groups: {except: :id}}
  end

  def update
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [EMPLOYEE_ROLE]

    @group = Group.find(params[:id])

    if @group.update(group_params)
      redirect_to @group.path
    else
      render json: {message: 'not updated', status: 400}
    end
  end

  def destroy
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [EMPLOYEE_ROLE]

    @group = Group.find(params[:id])
    @group.destroy
  end

  private

  def group_params
    params.require(:group).permit(:name)
  end
end