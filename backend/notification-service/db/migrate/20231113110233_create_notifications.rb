class CreateNotifications < ActiveRecord::Migration[7.1]
  def change
    create_table :notifications, id: :uuid, default: "gen_random_uuid()" do |t|
      t.datetime :event_time, null: false
      t.text :message, null: false
      t.datetime :display_from, null: false
      t.datetime :display_to, null: false
      t.uuid :recipient_id, null: false
    end
  end
end
