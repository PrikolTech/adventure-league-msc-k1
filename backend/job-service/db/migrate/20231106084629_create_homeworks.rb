class CreateHomeworks < ActiveRecord::Migration[7.1]
  def change
    create_table :homeworks, id: :uuid, default: "gen_random_uuid()" do |t|
      t.references :job, type: :uuid, null: false, foreign_key: true
      t.string :name
      t.string :description, null: true
      t.string :text, null: true

      t.timestamps
    end
  end
end
