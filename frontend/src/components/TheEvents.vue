<script setup>
import TheCalendar from '@/components/layouts/TheCalendar.vue';
import { onMounted, ref } from 'vue';
// дата в формате месяц день год
const events = ref([
    {
        format: "# Урок онлайн",
        title: "Стратегия развития анимации в отзывчивой дизайне",
        desc: "Курс : Дизайн система как смысл жизни",
        time: "20:00 - 22:00",
        date: '11.06.2023'
    },
    {
        format: "# Урок онлайн",
        title: "Стратегия развития анимации в отзывчивой дизайне",
        desc: "Курс : Дизайн система как смысл жизни",
        time: "20:00 - 22:00",
        date: '11.06.2023'
    },
    {
        format: "# Урок онлайн",
        title: "Стратегия развития анимации в отзывчивой дизайне",
        desc: "Курс : Дизайн система как смысл жизни",
        time: "20:00 - 22:00",
        date: '11.07.2023'
    },
    {
        format: "# Урок онлайн",
        title: "Стратегия развития анимации в отзывчивой дизайне",
        desc: "Курс : Дизайн система как смысл жизни",
        time: "20:00 - 22:00",
        date: '11.07.2023'
    },
    {
        format: "# Урок онлайн",
        title: "Стратегия развития анимации в отзывчивой дизайне",
        desc: "Курс : Дизайн система как смысл жизни",
        time: "20:00 - 22:00",
        date: '11.08.2023'
    },
    {
        format: "# Урок онлайн",
        title: "Стратегия развития анимации в отзывчивой дизайне",
        desc: "Курс : Дизайн система как смысл жизни",
        time: "20:00 - 22:00",
        date: '11.09.2023'
    },
    {
        format: "# Урок онлайн",
        title: "Стратегия развития анимации в отзывчивой дизайне",
        desc: "Курс : Дизайн система как смысл жизни",
        time: "20:00 - 22:00",
        date: '11.10.2023'
    },
    {
        format: "# Урок онлайн",
        title: "Стратегия развития анимации в отзывчивой дизайне",
        desc: "Курс : Дизайн система как смысл жизни",
        time: "20:00 - 22:00",
        date: '11.10.2023'
    },
])
const filteredEvents = ref([])
const datesWithEvents = ref([])

const getEvents = async() => {
    const response = [...events.value]

    filterEvents(formatDate(new Date()))

    response.forEach(el => {
        if (!datesWithEvents.value.includes(el.date)) {
            datesWithEvents.value.push(el.date);
        }
    });
    return events
}

const filterEvents = (date) => {
    filteredEvents.value = [
        ...events.value.filter(el => {
            return el.date ==formatDate(date)
        })
    ]
}

const formatDate = (date) => {
    const inputDate = new Date(date);
    const month = (inputDate.getMonth() + 1).toString().padStart(2, '0');
    const day = inputDate.getDate().toString().padStart(2, '0');
    const year = inputDate.getFullYear();

    const formattedDate = `${month}.${day}.${year}`;

    return formattedDate
}

onMounted(() => {
    getEvents()
})
</script>


<template>
    <div class="profile__events">
        <div class="profile__events-calendar">
            <the-calendar
                :dates="datesWithEvents"
                @datePick="(date) => filterEvents(date)"
            />
        </div>
        <div class="profile__events-list">
            <div class="profile__events-item"
                v-for="(event, index) of filteredEvents" :key="index"
            >
                <p class="profile__events-hashtag">
                    <!-- # Урок онлайн -->
                    {{ event.format }}
                </p>
                <p class="profile__events-title">
                    <!-- Стратегия развития анимации в отзывчивой дизайне -->
                    {{ event.title }}
                </p>
                <p class="profile__events-desc">
                    <!-- Курс : Дизайн система как смысл жизни -->
                    {{ event.desc }}
                </p>
                <p class="profile__events-time">
                    <!-- 20:00 - 22:00 -->
                    {{ event.time }}
                </p>
            </div>
        </div>
        <div class="text" v-if="filteredEvents.length === 0">
            Событий нет
        </div>
    </div>
</template>