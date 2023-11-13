class NotificationType::BaseNotification
  class << self
    def delta_from
      # Notification will be displayed 10 hours before lecture
      7.days
    end

    def delta_to
      # Notification will be displayed 1 hour after lecture
      7.days
    end

    def generate_message(event_time)
       "Событие начнется в #{event_time.strftime('%H:%M (%d-%m-%y)')}"
    end
  end
end