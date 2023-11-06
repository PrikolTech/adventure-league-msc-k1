<script setup>

</script>

<template>
    <main class="container notFound" style="display: flex; justify-content: center; align-items: center;">
        Страница не найдена
    </main>
</template>

<style lang="scss" scoped>
.notFound {

}

[dark=true] {
    & .notFound {
        color: var(--gray-50, #F9FAFB);
    }
}
</style>