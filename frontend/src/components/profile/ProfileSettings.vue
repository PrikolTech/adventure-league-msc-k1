<script setup>
import TheButton from '@/components/layouts/TheButton.vue';
import { useUser } from '@/stores/user'
const userStore = useUser()

</script>

<template>
    <div class="profile__settings profile__block">
        <div class="profile__block-header">
            <div class="profile__title title">
                Редактирование профиля
            </div>
            <div class="profile__header-line line"></div>
        </div>
        <div class="profile__settings-body">
            <div class="profile__settings-list">
                <div class="field">
                    <p>
                        Email
                    </p>
                    <div class="input-w">
                        <input
                            v-model="userStore.user.email"
                        >
                    </div>
                </div>
                <div class="field">
                    <p>
                        Номер телефона
                    </p>
                    <div class="input-w">
                        <input
                            v-model="userStore.user.phone"
                        >
                    </div>
                </div>
                <div class="field">
                    <p>
                        Telegram
                    </p>
                    <div class="input-w">
                        <input
                            v-model="userStore.user.telegram"
                        >
                    </div>
                </div>
            </div>
            <div class="profile__settings-btns">
                <the-button
                    :styles="['btn_red']"
                    :type="'button'"
                >
                    Сохранить
                </the-button>
                <the-button
                    :styles="['btn_grey']"
                    :type="'button'"
                >
                    Отмена
                </the-button>
            </div>
            <div class="profile__settings-footer">
                <p>
                    *Все ваши данные будут сразу удалены без возможности восстановления.
                </p>
                <the-button
                    :styles="['btn_red-border']"
                    :type="'button'"
                >
                    Удалить профиль
                </the-button>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.profile {

    &__settings-body {
        
    }

    &__settings-list {
        margin-bottom: 30px;
        display: flex;
        gap: 50px;
        flex-wrap: wrap;
        @media (max-width: 767px) {
            gap: 20px;
            margin-bottom: 20px;
        }
        & .field {
            flex: 0 0 100%;
            &:not(:last-child) {
                margin-bottom: 0px;
            }
        }
    }

    &__settings-btns {
        display: flex;
        gap: 30px;
        margin-bottom: 50px;
        @media (max-width: 767px) {
            margin-bottom: 25px;
        }
    }

    &__settings-footer {
        & p {
            color: #000;
            font-family: Roboto;
            font-size: 14px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 21px */
            display: block;
            margin-bottom: 12px;
        }
    }
}

[dark=true] {
    & .profile__settings-footer {
        & p {
            color: var(--gray-100, #F3F4F6);
        }
    }
}

</style>