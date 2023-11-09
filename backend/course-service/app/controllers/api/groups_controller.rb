class Api::GroupsController < ApplicationController
  def index
    @groups = Group.all
    render json: @groups
  end

  def show
    @group = Group.find(params[:id])

    render json: @group, include: {user_groups: {except: :id}}
  end

  def update
    @group = Group.find(params[:id])

    if @group.update(group_params)
      redirect_to @group.path
    else
      render json: {message: 'not updated', status: 400}
    end
  end

  def destroy
    @group = Group.find(params[:id])
    @group.destroy
  end

  private

  def group_params
    params.require(:group).permit(:name)
  end
end