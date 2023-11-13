<script setup>
import TheModal from "../layouts/TheModal.vue";

const props = defineProps({
    data: {
        type: Object,
        required: true,
    },
    formData: {
        type: Object,
        required: true,
    }
})
const emit = defineEmits(['sendForm'])

const namePopup = 'mainForm'


</script>

<template>
    <the-modal
        :name="namePopup"
        @send="emit('sendForm')"
    >
        <template v-slot:title>
            Заявка на программу
        </template>
        <div class="popup__list">
            <div class="field"
                v-for="(field, index) of props.data" :key="index"
            >
                <p>
                    {{ field.name }}
                </p>
                <div class="input-w">
                    <input
                        :placeholder="field.placeHolder"
                        :type="field.typeValue"
                        v-model="props.formData[field.nameInput]"
                        v-if="field.type === 'input'"
                    >
                    <textarea :placeholder="field.placeHolder" :type="field.typeValue" v-model="props.formData[field.nameInput]" v-if="field.type === 'textarea'"></textarea>
                </div>
            </div>
        </div>
        <div class="popup__agree agree" v-if="false">
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