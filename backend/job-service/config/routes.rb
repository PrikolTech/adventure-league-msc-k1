Rails.application.routes.draw do
  get "up" => "rails/health#show", as: :rails_health_check

  namespace :api do
    resources :jobs do
      resources :tests do
        resources :questions do
          resources :answers
        end

        resources :test_solutions, except: :update
      end

      resources :homeworks do
      end
    end
  end
end
