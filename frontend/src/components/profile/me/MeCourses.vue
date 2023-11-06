<script setup>

</script>


<template>
    <div class="me__courses">
        <div class="me__courses-list">
            <div class="me__courses-item"
                v-for="(mark, index) of [1,2,3]" :key="index"
            >
                <div class="me__courses-item-header">
                    <p>
                        Курс : Финансовая грамотность
                    </p>
                    <span>
                        Вы прошли половину обучения.
                    </span>
                </div>
                <div class="me__courses-item-content">
                    <div class="me__courses-data">
                        <p>
                            Просмотрено лекций :  
                        </p>
                        <span>
                            <b>
                                78 /
                            </b>
                            120
                        </span>
                    </div>
                    <div class="me__courses-data">
                        <p>
                            Выполнено ДЗ :
                        </p>
                        <span>
                            <b>
                                65 / 
                            </b>
                            100
                        </span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.me {

    &__courses {

    }

    &__courses-list {
        max-height: 600px;
        overflow: auto;
        border-radius: 20px;
    }

    &__courses-item {
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

    &__courses-item-header {
        & p {
            color: var(--gray-900, #111928);
            font-family: Inter;
            font-size: 18px;
            font-style: normal;
            font-weight: 500;
            line-height: 150%; /* 27px */
            @media (max-width: 767px) {
                font-size: 15px;
            }
        }
        & span {
            color: var(--gray-900, #111928);
            font-family: Roboto;
            font-size: 14px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 21px */
        }
    }

    &__courses-item-content {
        display: flex;
        gap: 40px;
    }

    &__courses-data {
        & p {
            color: var(--gray-500, #6B7280);
            font-family: Roboto;
            font-size: 12px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 18px */
        }
        & span {
            color: var(--unnamed, #7C8092);
            font-family: Roboto;
            font-size: 14px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%;
            @media (max-width: 767px) {
                font-size: 13px;
            }
            & b {
                color: var(--unnamed, #1F222E);
            }
        }
    }
}
[dark=true] {
    & .me__courses-item {
        background: #1A2537;
    }

    & .me__courses-item-header {
        & p {
            color: var(--gray-100, #F3F4F6);

        }
        & span {
            color: var(--gray-100, #F3F4F6);

        }
    }

    & .me__courses-data {
        & p {
            color: var(--gray-400, #9CA3AF);
        }
        & span {
            color: var(--gray-400, #9CA3AF);
            & b {
                color: var(--gray-100, #F3F4F6);
            }
        }
    }
}
</style>