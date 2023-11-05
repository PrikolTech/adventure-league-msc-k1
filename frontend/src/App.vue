<script setup>
import alertsList from '@/components/AlertsList.vue'
import TheHeader from '@/components/layouts/TheHeader.vue';
import TheFooter from '@/components/layouts/TheFooter.vue';
import { useUser } from '@/stores/user'
const user = useUser()
</script>

<template>
  <div class="wrapper" :dark="user.theme">
    <the-header/>
    <RouterView />
    <the-footer/>
    <alerts-list/>
  </div>
</template>

<style scoped>
.wrapper {
  transition: .2s;
  background: var(--var-body);
}
</style>
