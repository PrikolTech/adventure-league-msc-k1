<script setup>
import TheModal from "../layouts/TheModal.vue";
import { ref } from "vue";

const namePopup = 'mainForm'
let popupData = ref([
    {
        name: 'ФИО *',
        placeHolder: 'Иванов Иван Иванович',
        type: 'input',
        typeValye: 'text',
        value: '',
    },
    {
        name: 'Стаж *',
        placeHolder: '10',
        type: 'input',
        typeValye: 'number',
        value: '',
    },
    {
        name: 'Email *',
        placeHolder: 'pochta@mail.ru',
        type: 'input',
        typeValye: 'text',
        value: '',
    },
    {
        name: 'ФИО руководителя *',
        placeHolder: 'Иванов Иван Иванович',
        type: 'input',
        typeValye: 'text',
        value: '',
    },
    {
        name: 'Подразделение *',
        placeHolder: 'Placeholder',
        type: 'input',
        typeValye: 'text',
        value: '',
    },
    {
        name: 'Должность *',
        placeHolder: 'Менеджер',
        type: 'input',
        typeValye: 'text',
        value: '',
    },
    {
        name: 'Личные достижения в компании * (за последние 12 месяцев)',
        placeHolder: 'Напишите ваши достижения в компании...',
        type: 'textarea',
        typeValye: 'text',
        value: '',
    },
    {
        name: 'Мотивационное письмо *',
        placeHolder: 'Напишите почему Вы хотите учиться...',
        type: 'textarea',
        typeValye: 'text',
        value: '',
    },
])
</script>

<template>
    <the-modal
        :name="namePopup"
    >
        <template v-slot:title>
            Заявка на программу
        </template>
        <div class="popup__list">
            <div class="field"
                v-for="(field, index) of popupData" :key="index"
            >
                <p>
                    {{ field.name }}
                </p>
                <div class="input-w">
                    <input
                        :placeholder="field.placeHolder"
                        :type="field.typeValye"
                        v-model="field.value"
                        v-if="field.type === 'input'"
                    >
                    <textarea :placeholder="field.placeHolder" :type="field.typeValye" v-model="field.value" v-if="field.type === 'textarea'"></textarea>
                </div>
            </div>
        </div>
        <div class="popup__agree agree">
            <label class="b-contain">
                <span>Я согласен(а) на обработку персональных данных и с политикой обработки персональных данных</span>
                <input type="checkbox" />
                <div class="b-input"></div>
            </label>
        </div>
    </the-modal>
</template>

<style lang="scss" >

</style>