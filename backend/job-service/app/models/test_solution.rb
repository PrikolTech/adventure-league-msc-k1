class TestSolution < ApplicationRecord
  belongs_to :test
  has_many :solution_answers
  has_one :test_result

  def path
    "#{test.path}/test_solutions/#{id}"
  end

  def check_correctness
    # Returns mark in 0..5
    check_hash = {} # question_id: {correct: [], given: []}
    question_num = check_hash.keys.size

    answers_questions = {}
    solution_answers.each do |solution_answer|
      q_id = solution_answer.question.id
      answers_questions[q_id] = answers_questions.fetch(q_id, []).push(solution_answer.answer)
    end
    
    t_score = 0
    test.questions.each do |question|
      t_score += 1 if answers_questions[question.id].sort == question.correct_answers.sort
    end 
    
    (t_score.to_f / question_num * 5).round
  end
end
