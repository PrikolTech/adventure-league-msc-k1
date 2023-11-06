class ContentTypesController < ApplicationController
  def index
    @types = ContentType.all 
    render json: @types, only: :name
  end
end
