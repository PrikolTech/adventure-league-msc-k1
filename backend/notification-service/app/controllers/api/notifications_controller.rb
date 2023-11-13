class Api::NotificationsController < ApplicationController 
  EVENT_TYPES_TO_CLASS = {
    'LECTURE_START' => NotificationType::LectureStart,
    'HOMEWORK_RESULT' => NotificationType::HomeworkResult,
    'ADD_TO_COURSE' => NotificationType::AddToCourse,
    'BASE' => NotificationType::BaseNotification
  }
  
  def index
    @notifications = Notification.where(recipient_id: params[:sender_id])

    now = DateTime.now
    @notifications.where("display_to < ?", now).destroy_all
    @notifications = @notifications.where("display_from < ? AND display_to > ?", now, now)
  
    render json: {display: @notifications, all: Notification.all}
  end
  
  def create
    type_class = EVENT_TYPES_TO_CLASS[params[:event_type]]
    type_class = EVENT_TYPES_TO_CLASS['BASE'] if type_class.nil?

    event_time = params['event_time'].to_datetime 
    display_from = event_time - type_class.delta_from if event_time 
    display_to = event_time + type_class.delta_to if event_time
    message = type_class.generate_message(event_time)

    return render json: {message: 'incorrect display_to datetime', status: 400} if display_to < DateTime.now

    if @notification = Notification.create(
        recipient_id: params[:recipient_id],
        display_from: display_from,
        display_to: display_to,
        message: message,
        event_time: event_time
      )

      render json: @notification
    else
      render json: {message: 'not created', status: 400}
    end
  end
end
