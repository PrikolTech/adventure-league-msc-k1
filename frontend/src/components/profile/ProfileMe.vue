<script setup>
import TheButton from '@/components/layouts/TheButton.vue';
import TheAvatar from '@/components/layouts/TheAvatar.vue';
import { useUser } from '@/stores/user'
const userStore = useUser()

</script>

<template>
    <div class="profile__me profile__block">
        <div class="profile__me-body">
            <div class="profile__me-content">
                <div class="profile__block-header">
                    <div class="profile__title title">
                        Профиль
                    </div>
                    <div class="profile__header-line line"></div>
                </div>
                <div class="profile__me-header-info">
                    <the-avatar
                        :first_name="userStore.user.first_name"
                        :last_name="userStore.user.last_name"
                    />
                    <div class="profile__me-header-w">
                        <div class="profile__me-header">
                            <p>
                                <!-- Радкевич Наталья Владимировна -->
                                {{ userStore.user.first_name + ' ' + userStore.user.last_name }}
                            </p>
                            <the-button
                                :styles="['btn_grey']"
                                :type="'link'"
                                :link="'/profile/settings'"
                                style="gap: 10px;"
                            >
                                Редактировать
                                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="17" viewBox="0 0 16 17" fill="none">
                                    <path d="M7.33317 3.83333H3.99984C3.64622 3.83333 3.30708 3.97381 3.05703 4.22385C2.80698 4.4739 2.6665 4.81304 2.6665 5.16666V12.5C2.6665 12.8536 2.80698 13.1928 3.05703 13.4428C3.30708 13.6929 3.64622 13.8333 3.99984 13.8333H11.3332C11.6868 13.8333 12.0259 13.6929 12.276 13.4428C12.526 13.1928 12.6665 12.8536 12.6665 12.5V9.16666M11.7238 2.89066C11.8468 2.76332 11.994 2.66174 12.1566 2.59186C12.3193 2.52198 12.4943 2.4852 12.6713 2.48366C12.8483 2.48212 13.0239 2.51586 13.1878 2.5829C13.3516 2.64994 13.5005 2.74895 13.6257 2.87414C13.7509 2.99933 13.8499 3.1482 13.9169 3.31206C13.984 3.47592 14.0177 3.65149 14.0162 3.82853C14.0146 4.00557 13.9779 4.18053 13.908 4.3432C13.8381 4.50587 13.7365 4.653 13.6092 4.776L7.88517 10.5H5.99984V8.61466L11.7238 2.89066Z" stroke="#4B5563" stroke-linecap="round" stroke-linejoin="round"/>
                                </svg>
                            </the-button>
                        </div>
                        <div class="profile__me-groups">
                            <p>
                                Студент группы:
                            </p>
                            <div class="profile__me-groups-item">
                                ИДБ-20-09
                            </div>
                            <div class="profile__me-groups-item">
                                ИДБ-20-10
                            </div>
                        </div>
                    </div>
                </div>
                <div class="profile__me-tabs">
                    <div class="profile__me-tabs-header">
                        <router-link to="/profile/me/meCourses">
                            Курсы
                        </router-link>
                        <router-link to="/profile/me/meMarks">
                            Оценки
                        </router-link>
                        <router-link to="/profile/me/meComments">
                            Комментарии
                        </router-link>
                    </div>
                    <div class="profile__me-tabs-block">
                        <RouterView />
                    </div>
                </div>
            </div>
            <div class="profile__events-subtitle title">
                События
            </div>
        </div>
    </div>
</template>

<style lang="scss">
.profile__me-header-info {
    display: flex;
    gap: 20px;
    align-items: center;
    margin-bottom: 60px;
    @media (max-width: 1023px) {
        margin-bottom: 20px;
    }
    & .avatar {
        min-width: 80px;
        min-height: 80px;
        color: var(--gray-900, #111928);
        font-family: Inter;
        font-size: 30px;
        font-style: normal;
        font-weight: 500;
        line-height: 30px; /* 100% */
        @media (max-width: 767px) {
            display: none;
            min-width: 30px;
            min-height: 30px;
            font-size: 13px;
        }
    }
}
.profile {
    &__me-header-w {
        flex: 1;
    }

    &__me-body {
        display: flex;
        gap: 50px;
        justify-content: space-between;

        @media (max-width: 767px) {
            flex-wrap: wrap;
            gap: 20px;
        }
    }

    &__me-content {
        // max-width: 560px;
        flex: 1;
        @media (max-width: 767px) {
            max-width: none;
            flex: 0 0 100%;
        }
    }

    &__me-header {
        display: flex;
        column-gap: 50px;
        row-gap: 10px;
        flex-wrap: wrap;
        justify-content: space-between;
        margin-bottom: 15px;
        & p {
            color: var(--gray-900, #111928);
            font-family: Inter;
            font-size: 20px;
            font-style: normal;
            font-weight: 500;
            line-height: 150%; /* 30px */
            @media (max-width: 767px) {
                font-size: 16px;
            }
        }
        & a {
            font-family: Roboto;
            font-size: 14px;
            font-style: normal;
            font-weight: 500;
            line-height: 150%; /* 21px */
            padding: 9px 10px;
        }
    }

    &__me-groups {
        display: flex;
        align-items: center;
        gap: 12px;
        flex-wrap: wrap;
        @media (max-width: 1023px) {
        }
        @media (max-width: 767px) {
        }
        @media (max-width: 539px) {
        }
        & p {
            color: var(--gray-900, #111928);
            font-family: Roboto;
            font-size: 16px;
            font-style: normal;
            font-weight: 400;
            line-height: 167%; /* 26.72px */
        }
    }

    &__me-groups-item {
        border-radius: var(--rounded-xl, 12px);
        background: var(--gray-100, #F3F4F6);
        padding: 8px 10px;
        transition: .2s;
        font-family: Roboto;
        font-size: 14px;
        font-style: normal;
        font-weight: 400;
        line-height: 150%; /* 21px */
        @media (max-width: 767px) {
            font-size: 14px;
            padding: 4px 6px;
        }
    }

    &__me-tabs {

    }

    &__me-tabs-header {
        margin-bottom: 40px;
        display: flex;
        gap: 32px;
        width: fit-content;
        position: relative;
        @media (max-width: 1023px) {
            margin-bottom: 30px;
        }
        @media (max-width: 767px) {
            margin-bottom: 20px;
        }
        @media (max-width: 539px) {
            margin-bottom: 10px;
        }
        &::after {
            content: "";
            width: 100%;
            height: 1px;
            position: absolute;
            left: 0;
            bottom: 0;
            display: block;
            background: var(--var-line);
        }
        & a {
            color: var(--gray-500, #6B7280);
            font-family: Roboto;
            font-size: 16px;
            font-style: normal;
            font-weight: 500;
            line-height: 167%; /* 26.72px */
            padding-bottom: 15px;
            border-bottom: 1px solid transparent;
            z-index: 1;
            transition: .2s;
            &.router-link-active {
                color: var(--var-blue);
                border-bottom: 1px solid var(--var-blue);
            }
        }
    }

    &__me-tabs-block {
        
    }

    &__events {

    }
}

.profile {

    &__events {
        border-radius: 20px;
        background: var(--gray-200, #E5E7EB);
        padding: 30px;
        flex: 0 0 360px;
        height: fit-content;
        @media (max-width: 1199px) {
            flex: 0 0 320px;
        }
        @media (max-width: 1023px) {
            margin: 0 auto;
        }
        @media (max-width: 767px) {
            max-width: none;
            flex: 0 0 100%;
            padding: 20px;
        }
    }

    &__events-subtitle {
        display: none;
        @media (max-width: 767px) {
            display: block;
        }
    }

    &__events-calendar {
        margin-bottom: 20px;
        display: flex;
        justify-content: center;
    }

    &__events-list {
        display: flex;
        flex-direction: column;
        gap: 15px;
        border-radius: 12px;
        max-height: 700px;
        overflow: auto;
    }

    &__events-item {
        padding: 15px;
        background: #fff;
        border-radius: 12px;
        display: flex;
        flex-direction: column;
        gap: 8px;
    }

    &__events-hashtag {
        color: var(--gray-600, #4B5563);
        text-align: center;
        font-family: Roboto;
        font-size: 12px;
        font-style: normal;
        font-weight: 400;
        line-height: 150%; /* 18px */
        background: var(--gray-200, #E5E7EB);
        display: flex;
        width: fit-content;
        align-items: center;
        justify-content: center;
        padding: 4px 8px;
        border-radius: 8px;
    }

    &__events-title {
        color: #000;
        font-family: Roboto;
        font-size: 16px;
        font-style: normal;
        font-weight: 500;
        line-height: 150%; /* 24px */
    }

    &__events-desc {
        color: var(--gray-500, #6B7280);
        font-family: Roboto;
        font-size: 14px;
        font-style: normal;
        font-weight: 400;
        line-height: 150%; /* 21px */
    }

    &__events-time {
        color: var(--gray-800, #1F2A37);
        font-family: Roboto;
        font-size: 14px;
        font-style: normal;
        font-weight: 400;
        line-height: 150%; /* 21px */
    }
}

[dark=true] {
    & .profile__me-header p {
        color: var(--gray-50, #F9FAFB);
    }
    & .profile__me-header .btn {
        & path {
            stroke: #D1D5DB;
        }
    }
    & .profile__me-groups p {
        color: var(--gray-50, #F9FAFB);
    }

    & .profile__me-groups-item {
        color: var(--gray-50, #F9FAFB);
        background: var(--gray-800, #1F2A37);
    }

    & .profile__events {
        background: var(--gray-800, #1F2A37);
    }

    & .profile__events-item {
        background: var(--gray-700, #374151);
    }

    & .profile__events-hashtag {
        background: var(--gray-500, #6B7280);
        color: var(--gray-200, #E5E7EB);
    }
    
    & .profile__events-title {
        color: var(--gray-100, #F3F4F6);
    }

    & .profile__events-desc {
        color: var(--gray-400, #9CA3AF);
    }

    & .profile__events-time {
        color: var(--gray-100, #F3F4F6);
    }
}

</style>