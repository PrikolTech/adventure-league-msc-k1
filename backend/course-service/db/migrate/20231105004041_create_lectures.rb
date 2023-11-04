class CreateLectures < ActiveRecord::Migration[7.1]
  def change
    create_table :lectures, id: :uuid, default: "uuid_generate_v4()" do |t|
      t.string :name
      t.text :description
      t.datetime :start
      t.references :course, type: :uuid, null: false, foreign_key: true

      t.timestamps
    end
  end
end
