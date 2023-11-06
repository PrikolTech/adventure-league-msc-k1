Rails.application.routes.draw do
  get "up" => "rails/health#show", as: :rails_health_check

  namespace :api do
    resources :jobs do
      resources :tests do
        resources :questions do
          resources :answers, except: :show
        end
      end
    end
  end
end
