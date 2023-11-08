class Api::GroupsController < ActionController
  def index
    @groups = Group.all
    render json: @groups
  end

  def show
    @group = Group.find(params[:id])

    render json: @group, include: :user_group
  end

  def create
    @group = Group.create(group_params)

    if @group
      redirect_to @group.path
    else
      render json: {message: 'not created', status: 400}
    end
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

  def add_user
    group = Group.find(params[:id])
    user_id = params[:user_id]

    @u_g = UserGroup.new
    @u_g.group = group
    @u_g.user_id = user_id

    if @u_g
      render json: {message: 'ok', status: 201}
    else
      render json: {message: 'not created', status: 400}
    end
  end

  private

  def group_params
    params.require(:group).permit(:name)
  end
end