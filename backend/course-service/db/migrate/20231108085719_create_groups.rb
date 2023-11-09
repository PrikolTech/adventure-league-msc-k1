class CreateGroups < ActiveRecord::Migration[7.1]
  def change
    create_table :groups, id: :uuid, default: "gen_random_uuid()" do |t|
      t.string :name
      t.references :course, type: :uuid, null: false, foreign_key: true

      t.timestamps
    end
  end
end
