class CreateHomeworkResults < ActiveRecord::Migration[7.1]
  def change
    create_table :homework_results, id: :uuid, default: "gen_random_uuid()" do |t|
      t.uuid :user_id
      t.references :homework, type: :uuid, null: false, foreign_key: true
      t.integer :score

      t.timestamps
    end
  end
end
