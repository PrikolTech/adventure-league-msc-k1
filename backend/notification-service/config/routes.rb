Rails.application.routes.draw do
  namespace :api do
    resources :notifications, only: [:index] do
      collection do
        post 'create'
      end
      
    end
  end
end
