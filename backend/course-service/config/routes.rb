Rails.application.routes.draw do
  get 'api/course/education_forms', to: 'education_forms#index'
  get 'api/course/content_types', to: 'content_types#index'

  namespace :api do
    resources :course do
      resources :lecture do
        resources :content
      end
    end

    resources :groups do
      post 'add_user', to: 'groups#add_user' 
    end
  end
end
