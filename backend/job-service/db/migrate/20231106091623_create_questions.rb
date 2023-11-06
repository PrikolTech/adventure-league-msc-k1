class CreateQuestions < ActiveRecord::Migration[7.1]
  def change
    create_table :questions, id: :uuid, default: "gen_random_uuid()" do |t|
      t.references :test, type: :uuid, null: false, foreign_key: true
      t.text :body
      t.boolean :is_multiple_answer, default: false

      t.timestamps
    end
  end
end
