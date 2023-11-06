class CreateCourseTypes < ActiveRecord::Migration[7.1]
  def change
    create_table :course_types, id: :uuid, default: "gen_random_uuid()" do |t|
      t.string :name
    end
  end
end
