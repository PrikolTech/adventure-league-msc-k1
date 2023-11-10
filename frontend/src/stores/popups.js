import { ref } from 'vue'
import { defineStore } from 'pinia'
import { disableBodyScroll, enableBodyScroll } from 'body-scroll-lock';

export const usePopups = defineStore('popups', () => {
    const popups = ref(new Set([''])); // Create a new Set and initialize it with 'mainPopup'

    function disableScroll(block) {
        const element = document.querySelector(`.${block}`)
        disableBodyScroll(element)
        this.popups.add(block)
    }

    function enableScroll(block) {
        const element = document.querySelector(`.${block}`)
        this.popups.delete(block)
        enableBodyScroll(element)
    }

    return { popups, disableScroll, enableScroll }
})