# This file should ensure the existence of records required to run the application in every environment (production,
# development, test). The code here should be idempotent so that it can be executed at any point in every environment.
# The data can then be loaded with the bin/rails db:seed command (or created alongside the database with db:setup).
#
# Example:
#
#   ["Action", "Comedy", "Drama", "Horror"].each do |genre_name|
#     MovieGenre.find_or_create_by!(name: genre_name)
#   end

require 'json'

def load_seed(name)
  JSON.parse File.open(File.join(Rails.root, 'storage', 'seeds', "#{name}.json")).read
end

jobs_seed = load_seed('jobs')

jobs_seed.each do |job|
  Job.create(
    id: job['id'],
    lecture_id: job['lecture_id'],
    name: job['name'],
    deadline: job['deadline']
  )
end

puts '== Jobs seeding complete'

tests_seed = load_seed('tests')

tests_seed.each do |test|
  Test.create(
    id: test['id'],
    job_id: test['job_id'],
    name: test['name']
  )
end

puts '== Tests seeding complete'

questions_seed = load_seed('questions')

questions_seed.each do |question|
  Question.create(
    id: question['id'],
    test_id: question['test_id'],
    body: question['body'],
    is_multiple_answer: question['is_multiple_answer'],
  )
end

puts '== Questions seeding complete'

answers_seed = load_seed('answers')

answers_seed.each do |answer|
  Answer.create(
    id: answer['id'],
    question_id: answer['question_id'],
    body: answer['body'],
    is_correct: answer['is_correct']
  )
end

puts '== Answers seeding complete'
