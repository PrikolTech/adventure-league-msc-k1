Rails.application.routes.draw do
  get 'api/courses/education_forms', to: 'education_forms#index'
  get 'api/courses/content_types', to: 'content_types#index'

  namespace :api do
    resources :courses do
      post :add_user

      resources :lectures do
        resources :contents
      end
    end
  end
end
