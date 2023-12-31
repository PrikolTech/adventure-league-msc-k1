class CreateTestResults < ActiveRecord::Migration[7.1]
  def change
    create_table :test_results, id: :uuid, default: "gen_random_uuid()" do |t|
      t.references :test_solution, type: :uuid, null: false, foreign_key: true
      t.integer :score

      t.timestamps
    end
  end
end
