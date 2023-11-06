class CreatePeriods < ActiveRecord::Migration[7.1]
  def change
    create_table :periods, id: :uuid, default: "gen_random_uuid()" do |t|
      t.date :starts_at
      t.date :ends_at
    end
  end
end
