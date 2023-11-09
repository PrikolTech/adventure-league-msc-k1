<script setup>
import SunIcon from '../icons/SunIcon.vue';
import MonthIcon from '../icons/MonthIcon.vue';
import { useUser } from '@/stores/user'
const user = useUser()
const toggleTheme = () => {
    localStorage.setItem('theme', !user.theme)
}
</script>

<template>
    <div class="switcher">
        <input type="checkbox" id="switch" v-model="user.theme" @input="toggleTheme()"/>
        <label for="switch">Toggle</label>
        <sun-icon class="icon sun"/>
        <month-icon class="icon month"/>
    </div>
</template>

<style lang="scss" scoped>
.switcher {
    position: relative;
    width: fit-content;
    & .icon {
        position: absolute;
        top: 50%;
        transform: translate(0, -50%);
        z-index: 2;
        pointer-events: none;
        transition: .2s;
    }
    & .sun {
        right: 5px;
    }

    & .month {
        left: 5px;
        opacity: 0;
    }

    & input[type=checkbox]{
        height: 0;
        width: 0;
        visibility: hidden;
        position: absolute;
    }

    & label {
        cursor: pointer;
        text-indent: -9999px;
        width: 60px;
        height: 32px;
        background: #E5E7EB;
        display: block;
        border-radius: 100px;
        position: relative;
        transition: .2s;
    }

    & label:after {
        content: '';
        position: absolute;
        top: 5px;
        left: 5px;
        width: 22px;
        height: 22px;
        background: #fff;
        border-radius: 90px;
        transition: 0.3s;
    }

    & input:checked + label {
    }

    & input:checked + label:after {
        left: calc(100% - 5px);
        transform: translateX(-100%);
    }

    & label:active:after {
        width: 30px;
    }
}

[dark=true] {
    & .switcher {
        & label {
            background: var(--main-beerus, #374151);
        }
        & .sun {
            opacity: 0;
        }

        & .month {
            opacity: 1;
        }
    }
}
</style>