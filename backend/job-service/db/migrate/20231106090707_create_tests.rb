class CreateTests < ActiveRecord::Migration[7.1]
  def change
    create_table :tests, id: :uuid, default: "gen_random_uuid()" do |t|
      t.references :job, type: :uuid, null: false, foreign_key: true
      t.string :name

      t.timestamps
    end
  end
end
