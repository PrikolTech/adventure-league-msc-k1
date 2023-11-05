EDUCATION_FORMS = ['Онлайн', 'Очно']

EDUCATION_FORMS.each do |form|
  EducationForm.find_or_create_by!(name: form)
end

CONTENT_TYPES = ['video', 'file']

CONTENT_TYPES.each do |type|
  ContentType.find_or_create_by!(name: type)
end
