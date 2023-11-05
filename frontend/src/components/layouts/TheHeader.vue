<script setup>
import LogoIcon from '@/components/icons/LogoIcon.vue';
import switcherTheme from '@/components/layouts/SwitcherTheme.vue'
import { ref } from 'vue';
import TheButton from './TheButton.vue';

let headerMobileIsActive = ref(false)
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
            <div class="header__menu ">
                <div class="nav-icon toggle-menu"
                    :class="{open: headerMobileIsActive}"
                    @click="headerMobileIsActive = !headerMobileIsActive"
                >
                    <span></span>
                    <span></span>
                    <span></span>
                </div>
            </div>
        </div>
        <div class="header-m"
            :class="{active: headerMobileIsActive}"
        >
            <div class="header-m__content">
                <nav class="header-m__nav">
                    <router-link class="header-m__nav-item" to="/">
                        Обучение
                    </router-link>
                    <router-link class="header-m__nav-item" to="about">
                        Об университете
                    </router-link>
                    <router-link class="header-m__nav-item" to="events">
                        Мероприятия
                    </router-link>
                    <router-link class="header-m__nav-item" to="blog">
                        Блог
                    </router-link>
                    <router-link class="header-m__nav-item" to="contancts">
                        Контакты
                    </router-link>
                </nav>
                <div class="header-m__action">
                    <the-button
                        :styles="['btn_red-border']"
                        :type="'link'"
                        :link="'login'"
                        @click="headerMobileIsActive = false"
                    >
                        Войти
                    </the-button>
                    <div class="header__theme">
                        <switcher-theme/>
                    </div>
                </div>
                <div class="header-m__close toggle-menu"
                    @click="headerMobileIsActive = !headerMobileIsActive"
                >
                    <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 14 14" fill="none">
                        <path d="M1 1L13 13" stroke="#272727" stroke-width="1.5" stroke-linecap="round"/>
                        <path d="M13 1L1 13" stroke="#272727" stroke-width="1.5" stroke-linecap="round"/>
                    </svg>
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
    @media (max-width: 539px) {
        padding-top: 15px;
        padding-bottom: 15px;
        margin-bottom: 5px;
    }
    &__logo {
        display: flex;
        justify-content: center;
    }

    &__nav {
        display: flex;
        align-items: center;
        gap: 32px;
        @media (max-width: 1023px) {
            display: none;
        }
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
        @media (max-width: 1023px) {
            display: none;
        }
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


.nav-icon {
  width: 22px;
  height: 15px;
  position: relative;
  transform: rotate(0deg);
  transition: .5s ease-in-out;
  cursor: pointer;
}
.nav-icon span {
  display: block;
  position: absolute;
  height: 2.62px;
  width: 100%;
  background: #D1D5DB;
  border-radius: 10px;
  opacity: 1;
  left: 0;
  transform: rotate(0deg);
  transition: .25s ease-in-out;
}

.nav-icon span:nth-child(1) {
  top: 0px;
}
.nav-icon span:nth-child(2) {
  top: 6px;
}
.nav-icon span:nth-child(3) {
  top: 12px;
}
.nav-icon.open span {
  left: 7px;
}
.nav-icon.open span:nth-child(1) {
  top: 5px;
  transform: rotate(135deg);
  max-width: 30px;
  left: 0px;
}
.nav-icon.open span:nth-child(2) {
  opacity: 0;
  left: -60px;
}
.nav-icon.open span:nth-child(3) {
  top: 5px;
  transform: rotate(-135deg);
  max-width: 30px;
  left: 0px;
}
.header__menu {
  display: none;
  @media (max-width: 1023px) {
    display: block;
  }
}



.header-m {
      background: #F3F4F6;
      max-width: 290px;
      width: 63%;
      position: fixed;
      top: 0;
      left: 0;
      z-index: 10;
      left: -100%;
      transition: .4s;
      padding: 37px 15px 0px;
      height: 100vh;
      overflow: auto;
      &.active {
            left: 0;
      }
      &__content {
            display: flex;
            flex-direction: column;
            height: 100%;
            padding-bottom: 100px;
      }
      &__nav {
            display: flex;
            flex-direction: column;
            margin-bottom: 20px;
            gap: 10px;
      }
      &__nav-item {
        width: fit-content;
        position: relative;
        color: var(--var-grey);
        font-family: Roboto;
        font-size: 14px;
        font-style: normal;
        font-weight: 400;
        line-height: 150%;
        position: relative;
        padding-bottom: 4px;
        transition: 0.2s;
        &::after {
            content: "";
            position: absolute;
            left: 0;
            bottom: 0;
            display: block;
            width: 100%;
            height: 1px;
            background: var(--var-blue);
            border-radius: 50px;
            opacity: 0;
            transition: 0.2s;
        }
        &.router-link-active {
            &::after {
                opacity: 1;
            }
        }
      }
      & .header-m__action {
            display: flex;
            flex-direction: column-reverse;
            gap: 20px;
      }
}
.header-m__close {
      position: absolute;
      right: 10px;
      top: 13px;
      cursor: pointer;
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

    & .header-m {
        background: #2a3b56;
    }

    & .header-m__nav-item {
        color: #fff;
        &.router-link-active {
            &::after {
                opacity: 1;
                background: #fff;
            }
        }
    }

    & .header-m__close {
        & svg {
            & path {
                stroke: #fff;
            }
        }
    }
}

</style>