<script setup>
const props = defineProps({
    task: {
        type: Object,
        required: true
    }
})
</script>

<template>
    <div class="tasks__item">
        <p class="tasks__item-format">
            <!-- # ДЗ -->
            # {{ props.task.format }}
        </p>
        <p class="tasks__item-title">
            <!-- Сдать домашнее задание к уроку 3. Название урока. -->
            Курс : {{ props.task.title }}
        </p>
        <p class="tasks__item-desc">
            <!-- Курс : Финансовая грамотность -->
            {{ props.task.desc }}
        </p>
        <p class="tasks__item-time"
            v-if="props.task.status === 'Текущее'"
        >
            Дедлайн : 
            <b>
                <!-- Сегодня до 20:00 -->
                {{ props.task.time }}
            </b>
        </p>
        <p class="tasks__item-cheack"
            v-if="props.task.status === 'inspection'"
        >
            На проверке
        </p>
        <p class="tasks__item-mark"
            v-if="props.task.mark"
        >
            <span>
                Оценка: 
            </span>
            <b>
                <!-- 4 -->
                {{ props.task.mark }}
            </b>
        </p>
    </div>
</template>

<style lang="scss" scoped>
.tasks {

    &__item {
        border-radius: var(--rounded-xl, 12px);
        background: var(--gray-50, #F9FAFB);
        padding: 16px;
        flex: 0 0 calc(50% - 7.5px);
        @media (max-width: 539px) {
            flex: 0 0 100%;
        }
    }

    &__item-format {
        border-radius: 8px;
        background: var(--gray-200, #E5E7EB);
        padding: 4px 8px;
        color: var(--gray-600, #4B5563);
        font-family: Roboto;
        font-size: 12px;
        font-style: normal;
        font-weight: 400;
        line-height: 150%; /* 18px */
        margin-bottom: 8px;
        width: fit-content;
    }

    &__item-title {
        color: #000;
        font-family: Roboto;
        font-size: 16px;
        font-style: normal;
        font-weight: 500;
        line-height: 150%; /* 24px */
        margin-bottom: 2px;
    }

    &__item-desc {
        color: var(--gray-500, #6B7280);
        font-family: Roboto;
        font-size: 14px;
        font-style: normal;
        font-weight: 400;
        line-height: 150%; /* 21px */
        margin-bottom: 5px;
    }

    &__item-time {
        color: var(--gray-800, #1F2A37);
        font-family: Roboto;
        font-size: 14px;
        font-style: normal;
        font-weight: 400;
        line-height: 150%; /* 21px */
        & b {
            font-weight: 500;
        }
    }

    &__item-cheack {
        color: var(--gray-800, #1F2A37);
        font-family: Roboto;
        font-size: 14px;
        font-style: normal;
        font-weight: 500;
        line-height: 150%; /* 21px */
    }

    &__item-mark {
        display: flex;
        align-items: center;
        gap: 5px;
        & span {
            color: var(--gray-500, #6B7280);
            font-family: Roboto;
            font-size: 14px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 21px */
        }
        & b {
            min-width: 20px;
            min-height: 20px;
            display: flex;
            justify-content: center;
            align-items: center;
            color: var(--orange-500, #FF5A1F);
            font-family: Roboto;
            font-size: 12px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 18px */
            border-radius: 50%;
            background: #FFD695;
        }
    }
}

[dark=true] {
    & .tasks__item {
        border-radius: var(--rounded-xl, 12px);
        background: #1A2537;
    }
    & .tasks__item-format {
        color: var(--gray-50, #F9FAFB);
        background: var(--gray-700, #374151);
    }

    & .tasks__item-title {
        color: var(--gray-50, #F9FAFB);
    }

    & .tasks__item-desc {
        color: var(--gray-500, #6B7280);
    }

    & .tasks__item-time {
        color: var(--gray-50, #F9FAFB);
        & b {
            color: var(--gray-50, #F9FAFB);
        }
    }

    & .tasks__item-cheack {
        color: var(--gray-100, #F3F4F6);
    }

    & .tasks__item-mark {
        color: var(--gray-500, #6B7280);
        & b {
            color: var(--orange-400, #FF8A4C);
            background: #4c473e;
        }
    }
}
</style>