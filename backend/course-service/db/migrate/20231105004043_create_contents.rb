class CreateContents < ActiveRecord::Migration[7.1]
  def change
    create_table :contents, id: :uuid, default: "uuid_generate_v4()" do |t|
      t.string :url
      t.references :content_type, type: :uuid, null: false, foreign_key: true

      t.timestamps
    end
  end
end
