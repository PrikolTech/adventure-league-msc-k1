<script setup>

</script>

<template>
    <div class="me__marks">
        <div class="me__marks-header">

        </div>
        <div class="me__marks-list">
            <div class="me__marks-item"
                v-for="(mark, index) of [1,2,3]" :key="index"
            >
                <div class="me__marks-item-header">
                    <p>
                        # Экзамен
                    </p>
                    <span>
                        25 августа 2023
                    </span>
                </div>
                <p class="me__marks-item-title">
                    Подготовка к старту проекта
                </p>
                <p class="me__marks-item-course">
                    Курс : Agile, Scrum. 
                </p>
                <div class="me__marks-item-score">
                    <span>
                        5
                    </span>
                    <p>
                        Оценка преподавателя
                    </p>
                </div>
                <div class="me__marks-item-comment">
                    <textarea class="comment" readonly>Наталья, отличная работа! Молодец)</textarea>
                </div>
            </div>
        </div>
    </div>
</template>

<style lang="scss" >
.me {

    &__marks {

    }

    &__marks-header {

    }

    &__marks-list {
        max-height: 600px;
        overflow: auto;
        border-radius: 20px;
    }

    &__marks-item {
        border-radius: 20px;
        background: var(--gray-50, #F9FAFB);
        padding: 30px;
        display: flex;
        flex-direction: column;
        gap: 20px;
        @media (max-width: 767px) {
            padding: 15px;
            gap: 10px;
        }
        @media (max-width: 539px) {
        }
        &:not(:last-child) {
            margin-bottom: 30px;
            @media (max-width: 767px) {
                margin-bottom: 15px;
            }
        }
    }

    &__marks-item-header {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: 10px;
        flex-wrap: wrap;
        & p {
            border-radius: 8px;
            background: var(--gray-200, #E5E7EB);
            color: var(--gray-600, #4B5563);
            font-family: Roboto;
            font-size: 12px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 18px */
            padding: 4px 8px;
            display: flex;
            align-items: center;
            justify-content: center;
            text-align: center;
        }
        & span {
            color: var(--unnamed, #7C8092);
            font-family: Roboto;
            font-size: 14px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 21px */
        }
    }

    &__marks-item-title {
        color: var(--unnamed, #1F222E);
        font-size: 18px;
        font-style: normal;
        font-weight: 500;
        line-height: 150%; /* 27px */
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: 10px;
        @media (max-width: 767px) {
            font-size: 15px;
        }
        & span {
            color: var(--unnamed, #7C8092);
            font-family: Roboto;
            font-size: 14px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 21px */
            @media (max-width: 767px) {
                font-size: 13px;
            }
        }
    }

    &__marks-item-course {
        color: var(--gray-500, #6B7280);
        font-family: Roboto;
        font-size: 16px;
        font-style: normal;
        font-weight: 400;
        line-height: 167%; /* 26.72px */
    }

    &__marks-item-score {
        display: flex;
        gap: 12px;
        align-items: center;
        & span {
            display: flex;
            justify-content: center;
            align-items: center;
            min-width: 34px;
            min-height: 34px;
            background: #DDF7DA;
            border-radius: 50%;
            color: var(--green-500, var(--green-500, #0E9F6E));
            font-family: Roboto;
            font-size: 16px;
            font-style: normal;
            font-weight: 500;
            line-height: 150%; /* 24px */
        }
        & p {
            color: var(--unnamed, #1F222E);
            font-family: Roboto;
            font-size: 16px;
            font-style: normal;
            font-weight: 400;
            line-height: 167%; /* 26.72px */
        }
    }

    &__marks-item-comment {

    }
}


[dark=true] {
    & .me__marks-item {
        background: #1A2537;
    }

    & .me__marks-item-header {
        & p {
            color: var(--gray-200, #E5E7EB);
            background: var(--gray-600, #4B5563);
        }
        & span {
            color: var(--gray-500, #6B7280);
        }
    }

    & .me__marks-item-title {
        color: var(--gray-100, #F3F4F6);
    }

    & .me__marks-item-course {
        color: var(--gray-500, #6B7280);
    }

    & .me__marks-item-score {
        & p {
            color: var(--gray-200, #E5E7EB);
        }
    }
}
</style>