class CreateJobTypes < ActiveRecord::Migration[7.1]
  def change
    create_table :job_types do |t|
      t.string :name
    end
  end
end
