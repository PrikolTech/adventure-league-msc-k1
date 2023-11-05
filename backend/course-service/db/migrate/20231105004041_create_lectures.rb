class CreateLectures < ActiveRecord::Migration[7.1]
  def change
    create_table :lectures, id: :uuid, default: "gen_random_uuid()" do |t|
      t.string :name
      t.text :description
      t.datetime :starts_at
      t.references :course, type: :uuid, null: false, foreign_key: true

      t.timestamps
    end
  end
end
