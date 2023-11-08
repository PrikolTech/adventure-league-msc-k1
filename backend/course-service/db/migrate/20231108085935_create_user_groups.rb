class CreateUserGroups < ActiveRecord::Migration[7.1]
  def change
    create_table :user_groups do |t|
      t.references :group, type: :uuid, null: false, foreign_key: true
      t.uuid :user_id

      t.timestamps
    end
  end
end
