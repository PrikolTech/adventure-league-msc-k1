class CreateSolutionAnswers < ActiveRecord::Migration[7.1]
  def change
    create_table :solution_answers, id: :uuid, default: "gen_random_uuid()" do |t|
      t.references :test_solution, type: :uuid, null: false, foreign_key: true
      t.references :answer, type: :uuid, null: false, foreign_key: true

      t.timestamps
    end
  end
end
