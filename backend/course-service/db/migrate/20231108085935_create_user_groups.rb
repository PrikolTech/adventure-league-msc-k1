class CreateUserGroups < ActiveRecord::Migration[7.1]
  def change
    create_table :user_groups, id: :uuid, default: "gen_random_uuid()" do |t|
      t.references :group, type: :uuid, null: false, foreign_key: true
      t.uuid :user_id

      t.timestamps
    end
  end
end
