<script setup>
import alertsList from '@/components/AlertsList.vue'
import TheHeader from '@/components/layouts/TheHeader.vue';
import { useUser } from '@/stores/user'
const user = useUser()
</script>

<template>
  <div class="wrapper" :class="{dark: user.theme}">
    <the-header/>
    <RouterView />
    <alerts-list/>
  </div>
</template>

<style scoped>
</style>
