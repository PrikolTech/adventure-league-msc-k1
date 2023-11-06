class CreateJobs < ActiveRecord::Migration[7.1]
  def change
    create_table :jobs, id: :uuid, default: "gen_random_uuid()" do |t|
      t.uuid :lecture_id
      t.string :name
      t.text :description
      t.datetime :deadline, null: true

      t.timestamps
    end
  end
end
