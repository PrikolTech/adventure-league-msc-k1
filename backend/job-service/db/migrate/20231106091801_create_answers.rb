class CreateAnswers < ActiveRecord::Migration[7.1]
  def change
    create_table :answers, id: :uuid, default: "gen_random_uuid()" do |t|
      t.references :question, type: :uuid, null: false, foreign_key: true
      t.boolean :is_correct, default: true
      t.text :body

      t.timestamps
    end
  end
end
