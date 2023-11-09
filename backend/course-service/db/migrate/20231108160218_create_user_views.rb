class CreateUserViews < ActiveRecord::Migration[7.1]
  def change
    create_table :user_views, id: :uuid, default: "gen_random_uuid()" do |t|
      t.references :lecture, type: :uuid, null: false, foreign_key: true
      t.uuid :user_id

      t.timestamps
    end
  end
end
