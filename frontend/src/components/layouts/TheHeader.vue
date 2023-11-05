<script setup>
import LogoIcon from '@/components/icons/LogoIcon.vue';
import SearchIcon from '@/components/icons/SearchIcon.vue';
import closeIcon from '@/components/icons/CloseIcon.vue'
import switcherTheme from '@/components/layouts/SwitcherTheme.vue'
import TheButton from './TheButton.vue';
</script>

<template>
    <header class="header-w">
        <div class="header container">
            <router-link class="header__logo logo" to="/">
                <logo-icon/>
            </router-link>
            <nav class="header__nav">
                <router-link class="header__nav-item" to="/">Обучение</router-link>
                <router-link class="header__nav-item" to="about">Об университете</router-link>
                <router-link class="header__nav-item" to="events">Мероприятия</router-link>
                <router-link class="header__nav-item" to="blog">Блог</router-link>
                <router-link class="header__nav-item" to="contancts">Контакты</router-link>
            </nav>
            <div class="header__action">
                <div class="header__search">
                    <input>
                    <search-icon
                        class="search-big-icon"
                    />
                    <close-icon
                        class="close-icon"
                    />
                    <svg class="search-mini-icon" xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 12 12" fill="none">
                        <path d="M4.80295 0C3.85302 0 2.92442 0.281688 2.13458 0.809443C1.34473 1.3372 0.729129 2.08732 0.365605 2.96494C0.00208103 3.84256 -0.0930334 4.80828 0.0922895 5.73996C0.277612 6.67164 0.735049 7.52744 1.40675 8.19915C2.07846 8.87085 2.93426 9.32829 3.86594 9.51361C4.79762 9.69894 5.76334 9.60382 6.64096 9.2403C7.51859 8.87677 8.2687 8.26117 8.79646 7.47133C9.32421 6.68148 9.6059 5.75288 9.6059 4.80295C9.60447 3.52957 9.09799 2.30875 8.19757 1.40833C7.29715 0.507913 6.07633 0.00143014 4.80295 0ZM4.80295 8.40516C4.0905 8.40516 3.39405 8.1939 2.80167 7.79808C2.20929 7.40226 1.74758 6.83967 1.47494 6.18146C1.2023 5.52324 1.13096 4.79895 1.26996 4.10019C1.40895 3.40143 1.75202 2.75958 2.2558 2.2558C2.75958 1.75202 3.40143 1.40895 4.1002 1.26995C4.79896 1.13096 5.52324 1.2023 6.18146 1.47494C6.83968 1.74758 7.40227 2.20929 7.79808 2.80167C8.1939 3.39405 8.40516 4.0905 8.40516 4.80295C8.40421 5.75802 8.02439 6.67371 7.34905 7.34905C6.67371 8.02438 5.75803 8.40421 4.80295 8.40516Z" fill="#1F2A37"/>
                        <path d="M11.8315 10.9825L9.42999 8.58107C9.31676 8.47171 9.16511 8.41119 9.00769 8.41256C8.85028 8.41393 8.6997 8.47707 8.58839 8.58838C8.47707 8.6997 8.41393 8.85028 8.41256 9.00769C8.4112 9.16511 8.47171 9.31676 8.58107 9.42999L10.9825 11.8315C11.0958 11.9408 11.2474 12.0013 11.4048 12C11.5623 11.9986 11.7128 11.9355 11.8242 11.8242C11.9355 11.7128 11.9986 11.5623 12 11.4048C12.0013 11.2474 11.9408 11.0958 11.8315 10.9825Z" fill="#1F2A37"/>
                    </svg>
                </div>
                <div class="header__theme">
                    <switcher-theme/>
                </div>
                <div class="header__auth">
                    <the-button
                        :styles="['btn_red-border']"
                        :type="'link'"
                        :link="'login'"
                    >
                        Войти
                    </the-button>
                </div>
            </div>
        </div>
    </header>
</template>

<style lang="scss" scoped>
.header-w {
    position: sticky;
    top: 0;
    z-index: 10;
    background: var(--var-body);
}
.header {
    display: flex;
    align-items: center;
    gap: 20px;
    justify-content: space-between;
    padding-top: 20px;
    padding-bottom: 20px;
    margin-bottom: 35px;
    &__logo {
    }

    &__nav {
        display: flex;
        align-items: center;
        gap: 32px;
        
    }

    &__nav-item {
        color: var(--var-grey);
        font-family: Roboto;
        font-size: 14px;
        font-style: normal;
        font-weight: 400;
        line-height: 150%; /* 21px */
        position: relative;
        padding-bottom: 4px;
        transition: .2s;
        &::after {
            content: "";
            position: absolute;
            left: 4px;
            bottom: 0;
            display: block;
            width: calc(100% - 8px);
            height: 1px;
            background: var(--var-blue);
            border-radius: 50px;
            opacity: 0;
            transition: .2s;
        }
        &.router-link-active {
            color: var(--var-blue);
            &::after {
                opacity: 1;
            }
        }
    }
    
    &__search {
        position: relative;
        & input {
            border-radius: var(--rounded-xl, 12px);
            border: 1px solid var(--main-beerus, #E2E2E2);
            background: var(--main-gohan, #FFF);
            width: 100%;
            padding: 7px 16px 7px 25px;
        }
        & .search-big-icon {
            position: absolute;
            display: none;
        }
        & .search-mini-icon {
            position: absolute;
            left: 10px;
            top: 50%;
            transform: translate(0, -50%);
        }
        & .close-icon {
            position: absolute;
            right: 10px;
            top: 50%;
            transform: translate(0, -50%);
        }
    }

    &__action {
        display: flex;
        align-items: center;
        gap: 23px;
    }

    &__theme {
    }

    &__auth {
        padding-left: 23px;
        position: relative;
        &::after {
            content: "";
            position: absolute;
            left: 0;
            top: 5px;
            height: 34px;
            background: #D1D5DB;
            display: block;
            width: 1.5px;
            border-radius: 20px;
        }
    }
}

[dark=true] {
    & .header__nav-item {
        &.router-link-active {
            color: #fff;
            &::after {
                opacity: 1;
                background: #fff;
            }
        }
    }
    & .header__auth {
        &::after {
            background: #4B5563;
        }
    }
}

</style>