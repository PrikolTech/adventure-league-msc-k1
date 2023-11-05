<script setup>
import TheButton from '@/components/layouts/TheButton.vue';

</script>

<template>
    <main>
      <div class="login container">
        <div class="login__header">
            <div class="login__header-line"></div>
            <div class="login__title">
                Авторизация
            </div>
            <div class="login__header-line"></div>
        </div>
        <div class="login__form">
            <div class="field">
                <p>
                    Email *
                </p>
                <div class="input-w">
                    <input
                        placeholder=""
                    >
                </div>
            </div>
            <div class="field">
                <p>
                    Пароль
                </p>
                <div class="input-w">
                    <input
                        placeholder=""
                    >
                </div>
            </div>
            <div class="login__agree agree">
                <label class="b-contain">
                    <span>Запомните меня</span>
                    <input type="checkbox" />
                    <div class="b-input"></div>
                </label>
            </div>
            <the-button
                class="login__btn"
                :styles="['btn_red']"
                :type="'button'"
            >
                Войти
            </the-button>
        </div>
      </div>
    </main>
</template>

<style lang="scss" >
.login {
    max-width: 390px;
    &__header {
        display: flex;
        align-items: center;
        margin-bottom: 40px;
        @media (max-width: 539px) {
            margin-bottom: 10px;
        }
    }

    &__header-line {
        flex: 1;
    }

    &__title {
        color: var(--var-grey-lite-font);
    }

    &__form {
    }

    &__agree {
    }

    &__btn {
        margin-top: 30px;
        @media (max-width: 539px) {
            width: 100%;
            margin-top: 20px;
        }
        & .btn {
            width: 100%;
        }
    }
}


</style>