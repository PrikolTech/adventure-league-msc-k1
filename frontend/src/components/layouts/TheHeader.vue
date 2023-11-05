<script setup>
import LogoIcon from '@/components/icons/LogoIcon.vue';
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
    transition: .2s;
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