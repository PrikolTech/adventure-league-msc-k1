class TestSolution < ApplicationRecord
  belongs_to :test
  has_many :solution_answers

  def path
    "#{test.path}/#{id}"
  end

  def check_correctness
    check_hash = {} # question_id: {correct: [], given: []}

    answers_questions = {}
    solution_answers.each do |solution_answer|
      q_id = solution_answer.question.id
      answers_questions[q_id] = answers_questions.fetch(q_id, []).push(solution_answer.answer)
    end
    
    t_score = 0
    test.questions.each do |question|
      t_score += answers_questions[question.id].sort == question.answers.to_a.sort
    end
    
    t_score
  end
end
