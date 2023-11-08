# EDUCATION_FORMS = ['Онлайн', 'Очно']

# EDUCATION_FORMS.each do |form|
#   EducationForm.find_or_create_by!(name: form)
# end

# CONTENT_TYPES = ['video', 'file']

# CONTENT_TYPES.each do |type|
#   ContentType.find_or_create_by!(name: type)
# end


require 'json'

def load_seed(name)
  JSON.parse File.open(File.join(Rails.root, 'db', 'seeds', "#{name}.json")).read
end

course_types_seeds = load_seed('course_types')

course_types_seeds.each do |course_type|
  CourseType.create(
    id: course_type['id'],
    name: course_type['name']
  )
end

puts '== Course Types seeding complete'

education_forms_seed = load_seed('education_forms')

education_forms_seed.each do |education_form|
  EducationForm.create(
    id: education_form['id'],
    name: education_form['name']
  )
end

puts '== Education Forms seeding complete'

periods_seed = load_seed('periods')

periods_seed.each do |period|
  Period.create(
    id: period['id'],
    starts_at: period['starts_at'],
    ends_at: period['ends_at']
  )
end

puts '== Periods seeding complete'

lectures_seed = load_seed('lectures')

lectures_seed.each do |lecture|
  Lecture.create(
    id: lecture['id'],
    name: lecture['name'],
    description: lecture['description'],
    starts_at: lecture['starts_at'],
    course_id: lecture['course_id']
  )
end

puts '== Lectures seeding complete'

courses_seed = load_seed('courses')

courses_seed.each do |course|
  Course.create(
    id: course['id'],
    name: course['name'],
    description: course['description'],
    advantages: course['advantages'],
    price: course['price'],
    course_type_id: course['course_type_id'],
    period_id: course['period_id'],
    education_form_id: course['education_form_id'],
    prefix: course['prefix'],
  )
end

puts '== Courses seeding complete'

