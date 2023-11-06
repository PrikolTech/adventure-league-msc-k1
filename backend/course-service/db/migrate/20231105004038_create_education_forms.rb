class CreateEducationForms < ActiveRecord::Migration[7.1]
  def change
    create_table :education_forms, id: :uuid, default: "gen_random_uuid()" do |t|
      t.string :name
    end
  end
end
