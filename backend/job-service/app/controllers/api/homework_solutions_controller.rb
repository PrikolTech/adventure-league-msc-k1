class Api::HomeworkSolutionsController < ApplicationController
  EXTERNAL_ERROR = 'file service error'.freeze
  
  def index
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TEACHER_ROLE, STUDENT_ROLE]
    
    @solutions = []
    @homework = Homework.find(params[:homework_id])
    if has_permission? [TEACHER_ROLE]
      @solutions = HomeworkSolution.where(homework_id: @homework.id)
      if params.include? :user_id
        @solutions = @solutions.where(user_id: params[:user_id])
      end
    else
      @solutions = HomeworkSolution.where(homework_id: @homework.id, user_id: params[:sender_id])
    end

    render json: {solutions: @solutions, include: :homework_result, homework: @homework}
  end
  
  def show
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TEACHER_ROLE, STUDENT_ROLE]

    if has_permission? [TEACHER_ROLE]
      @solution = HomeworkSolution.find(params[:id])
    else
      @solution = HomeworkSolution.find_by(id: params[:id], user_id: params[:sender_id])
    end

    render json: @solution, include: :homework_result
  end

  def create
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [STUDENT_ROLE]
    return render json: {message: 'не ауф', status: 401} unless params.include? :sender_id

    @solution = HomeworkSolution.new(user_id: params[:sender_id])
    
    if @solution
      begin
        url = Intercommunication::FileService.send_file(params[:file])
        raise StandardError.new if url.nil? 
      rescue => exception
        return render json: {message: EXTERNAL_ERROR, status: 500}
      end

      @solution.file_url = url
      @solution.homework_id = params[:homework_id]
      @solution.save
      
      render json: @solution, include: :homework_result
    else
      render json: {message: 'not created', status: 400}
    end
  end

  def update
    return render json: {message: 'dont knock this door', status: 403}

    @solution = HomeworkSolution.find_by(id: params[:id], user_id: params[:sender_id])
    return render json: {status: 403, message: NO_PERMISSION_ERROR} if @solution.nil?

    if @solution.update(homework_solution_params)
      redirect_to @solution.path
    else
      render json: {message: 'not updated', status: 400}
    end
  end

  def destroy
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [STUDENT_ROLE]

    @solution = HomeworkSolution.find_by(id: params[:id], user_id: params[:sender_id])
    return render json: {status: 403, message: NO_PERMISSION_ERROR} if @solution.nil?

    HomeworkResult.where(solution_id: @solution.id)

    @solution.destroy
  end

  def result
    return render json: {status: 403, message: NO_PERMISSION_ERROR} unless has_permission? [TEACHER_ROLE]
    
    @solution = HomeworkSolution.find(params[:homework_solution_id])
    HomeworkResult.create(score: params[:score], homework_solution_id: @solution.id)

    render json: @solution, include: :homework_result
  end

  private
end
