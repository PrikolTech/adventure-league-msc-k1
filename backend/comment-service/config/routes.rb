Rails.application.routes.draw do
  post "api/comments/:target_type/", to: "comments#create"
  get  "api/comments/:target_type/:target_id", to: "comments#index"
  # root "posts#index"
end
