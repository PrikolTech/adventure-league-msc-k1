class CreateCourses < ActiveRecord::Migration[7.1]
  def change
    create_table :courses, id: :uuid, default: "uuid_generate_v4()" do |t|
      t.string :name, null: false
      t.text :description, null: true
      t.float :price, null: true
      t.references :course_type, type: :uuid, null: false, foreign_key: true
      t.references :period, type: :uuid, null: false, foreign_key: true
      t.references :education_form, type: :uuid, null: false, foreign_key: true

      t.timestamps
    end
  end
end
