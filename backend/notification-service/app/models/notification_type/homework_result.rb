class NotificationType::HomeworkResult < NotificationType::BaseNotification
  class << self
    def delta_from
      0
    end

    def delta_to
      # Notification will be displayed 7 days after lecture
      7.days
    end

    def generate_message(event_time)
       "Твою работу оценили #{event_time.strftime('%d-%m-%y')} числа. Быстрее проверь результаты!"
    end
  end
end