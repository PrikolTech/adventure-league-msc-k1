import { ref } from 'vue'
import { defineStore } from 'pinia'
import { disableBodyScroll, enableBodyScroll } from 'body-scroll-lock';

export const usePopups = defineStore('popups', () => {
    const popups = ref(new Set(['mainPopup'])); // Create a new Set and initialize it with 'mainPopup'

    function disableScroll() {
        disableBodyScroll
    }

    function enableScroll() {
        enableBodyScroll()
    }

    return { popups, disableScroll, enableScroll }
})