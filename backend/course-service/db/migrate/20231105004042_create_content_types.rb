class CreateContentTypes < ActiveRecord::Migration[7.1]
  TYPES = ['video', 'file']

  def change
    create_table :content_types, id: :uuid, default: "gen_random_uuid()" do |t|
      t.string :name
    end
  end

  def up
    TYPES.each do |type|
      ContentType.create(name: type)
    end
  end
end
