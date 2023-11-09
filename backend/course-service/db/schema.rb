# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# This file is the source Rails uses to define your schema when running `bin/rails
# db:schema:load`. When creating a new database, `bin/rails db:schema:load` tends to
# be faster and is potentially less error prone than running all of your
# migrations from scratch. Old migrations may fail to apply correctly if those
# migrations use external dependencies or application code.
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema[7.1].define(version: 2023_11_08_160218) do
  # These are extensions that must be enabled in order to support this database
  enable_extension "plpgsql"

  create_table "content_types", id: :uuid, default: -> { "gen_random_uuid()" }, force: :cascade do |t|
    t.string "name"
  end

  create_table "contents", id: :uuid, default: -> { "gen_random_uuid()" }, force: :cascade do |t|
    t.string "url"
    t.uuid "content_type_id", null: false
    t.uuid "lecture_id", null: false
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["content_type_id"], name: "index_contents_on_content_type_id"
    t.index ["lecture_id"], name: "index_contents_on_lecture_id"
  end

  create_table "course_types", id: :uuid, default: -> { "gen_random_uuid()" }, force: :cascade do |t|
    t.string "name"
  end

  create_table "courses", id: :uuid, default: -> { "gen_random_uuid()" }, force: :cascade do |t|
    t.string "name", null: false
    t.text "description"
    t.text "advantages"
    t.float "price"
    t.uuid "course_type_id", null: false
    t.uuid "period_id", null: false
    t.uuid "education_form_id", null: false
    t.string "tg_link"
    t.string "prefix", null: false
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["course_type_id"], name: "index_courses_on_course_type_id"
    t.index ["education_form_id"], name: "index_courses_on_education_form_id"
    t.index ["period_id"], name: "index_courses_on_period_id"
  end

  create_table "education_forms", id: :uuid, default: -> { "gen_random_uuid()" }, force: :cascade do |t|
    t.string "name"
  end

  create_table "groups", id: :uuid, default: -> { "gen_random_uuid()" }, force: :cascade do |t|
    t.string "name"
    t.uuid "course_id", null: false
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["course_id"], name: "index_groups_on_course_id"
  end

  create_table "lectures", id: :uuid, default: -> { "gen_random_uuid()" }, force: :cascade do |t|
    t.string "name"
    t.text "description"
    t.datetime "starts_at"
    t.uuid "course_id", null: false
    t.boolean "is_hidden", default: false
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["course_id"], name: "index_lectures_on_course_id"
  end

  create_table "periods", id: :uuid, default: -> { "gen_random_uuid()" }, force: :cascade do |t|
    t.date "starts_at"
    t.date "ends_at"
  end

  create_table "user_groups", id: :uuid, default: -> { "gen_random_uuid()" }, force: :cascade do |t|
    t.uuid "group_id", null: false
    t.uuid "user_id"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["group_id"], name: "index_user_groups_on_group_id"
  end

  create_table "user_views", id: :uuid, default: -> { "gen_random_uuid()" }, force: :cascade do |t|
    t.uuid "lecture_id", null: false
    t.uuid "user_id"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["lecture_id"], name: "index_user_views_on_lecture_id"
  end

  add_foreign_key "contents", "content_types"
  add_foreign_key "contents", "lectures"
  add_foreign_key "courses", "course_types"
  add_foreign_key "courses", "education_forms"
  add_foreign_key "courses", "periods"
  add_foreign_key "groups", "courses"
  add_foreign_key "lectures", "courses"
  add_foreign_key "user_groups", "groups"
  add_foreign_key "user_views", "lectures"
end
