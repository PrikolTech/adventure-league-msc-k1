Rails.application.routes.draw do
  get "up" => "rails/health#show", as: :rails_health_check

  get 'api/course_types', to: 'course_types#index'
  get 'api/education_forms', to: 'education_forms#index'

  get 'api/courses', to: 'courses#index'
  get 'api/courses/:id', to: 'courses#show'
  post 'api/courses/create', to: 'courses#create'

  get 'api/courses/:course_id/lectures', to: 'lectures#index'
  post 'api/courses/:course_id/lectures/create', to: 'lectures#create'

  get 'api/lectures/:id', to: 'lectures#show'
  post 'api/lectures/:id/contents/create', to: 'contents#create'
end
