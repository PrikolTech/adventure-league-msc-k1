class CreateHomeworkSolutions < ActiveRecord::Migration[7.1]
  def change
    create_table :homework_solutions, id: :uuid, default: "gen_random_uuid()" do |t|
      t.references :homework, type: :uuid, null: false, foreign_key: true 
      t.uuid :user_id
      t.string :file_url

      t.timestamps
    end
  end
end
