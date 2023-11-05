class EducationFormsController < ApplicationController
  def index
    @forms = EducationForm.all
    render json: @forms, only: :name
  end
end
