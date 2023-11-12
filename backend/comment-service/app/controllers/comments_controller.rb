class CommentsController < ApplicationController
  def index
    target_type_id = TargetType.find_by(name: params[:target_type]).id
    @comments = Comment.where(target_id: params[:target_id], target_type_id: target_type_id)
  
    render json: @comments, except: :target_type_id
  end

  def create
    if @comment = Comment.create(
        target_type_id: TargetType.find_by(name: params[:target_type]).id,
        body: params[:body],
        user_id: params[:user_id],
        target_id: params[:target_id]
      )

      render json: @comment, except: :target_type_id
    else
      render json: {message: 'not created', status: 400}
    end
  end

  def destroy
    @comment = Comment.find(params[:id])
    @comment.destroy
  end
end