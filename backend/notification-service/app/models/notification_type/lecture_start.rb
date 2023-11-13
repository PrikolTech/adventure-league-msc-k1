class NotificationType::LectureStart < NotificationType::BaseNotification
  class << self
    def delta_from
      # Notification will be displayed 10 hours before lecture
      10.hours
    end

    def delta_to
      # Notification will be displayed 1 hour after lecture
      1.hours
    end

    def generate_message(event_time)
       "Не забудь про лекцию. Она начнется #{event_time.strftime('%H:%M')}"
    end
  end
end