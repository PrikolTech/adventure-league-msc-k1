class CreateTestSolutions < ActiveRecord::Migration[7.1]
  def change
    create_table :test_solutions, id: :uuid, default: "gen_random_uuid()" do |t|
      t.references :test, type: :uuid, null: false, foreign_key: true
      t.uuid :user_id

      t.timestamps
    end
  end
end
