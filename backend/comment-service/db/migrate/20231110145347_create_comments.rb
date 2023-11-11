class CreateComments < ActiveRecord::Migration[7.1]
  def change
    create_table :comments, id: :uuid, default: "gen_random_uuid()" do |t|
      t.text :body, null: false
      t.references :target_type, null: false, foreign_key: true
      t.uuid :user_id, null: false
      t.uuid :target_id, null: false

      t.timestamps
    end
  end
end
