Rails.application.routes.draw do
  get "up" => "rails/health#show", as: :rails_health_check

  get '/', to: 'courses#index'
end
