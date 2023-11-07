<script setup>
const props = defineProps({
    links: {
        type: Array,
        required: true
    }
})
</script>

<template>
    <div class="navigation">
        <router-link
            v-for="(link, index) of props.links" :key="index"
            :to="link.href"
        >
            {{ link.text }}
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="21" viewBox="0 0 20 21" fill="none">
                <path d="M7.7895 16.3333C7.57022 16.3333 7.35587 16.2649 7.17356 16.1369C6.99125 16.0088 6.84916 15.8269 6.76525 15.614C6.68134 15.4011 6.65938 15.1668 6.70215 14.9408C6.74492 14.7148 6.85049 14.5072 7.00552 14.3443L10.6571 10.5071L7.00552 6.67003C6.89961 6.56254 6.81513 6.43396 6.75702 6.2918C6.6989 6.14964 6.66831 5.99674 6.66703 5.84202C6.66575 5.6873 6.69381 5.53386 6.74957 5.39066C6.80532 5.24745 6.88766 5.11735 6.99178 5.00795C7.09589 4.89854 7.2197 4.81202 7.35598 4.75343C7.49225 4.69484 7.63827 4.66536 7.78551 4.6667C7.93274 4.66804 8.07825 4.70019 8.21354 4.76126C8.34883 4.82233 8.47119 4.9111 8.57348 5.02239L13.009 9.68333C13.2169 9.90184 13.3337 10.1982 13.3337 10.5071C13.3337 10.8161 13.2169 11.1125 13.009 11.331L8.57348 15.9919C8.36557 16.2104 8.08357 16.3333 7.7895 16.3333Z" fill="#9CA3AF"/>
            </svg>
        </router-link>
    </div>
</template>

<style lang="scss" scoped>
.navigation {
    display: flex;
    flex-wrap: wrap;
    gap: 14px;
    & a {
        display: flex;
        align-items: center;
        gap: 16px;
        color: var(--gray-500, #6B7280);
        font-family: Roboto;
        font-size: 14px;
        font-style: normal;
        font-weight: 400;
        line-height: 150%; /* 21px */
        &:last-child {
            color: var(--gray-800, #1F2A37);
            & svg {
                display: none;
            }
        }
    }
}

[dark=true] {
    .navigation {
    & a {
        color: var(--gray-500, #6B7280);
        & svg {
            
        }
        &:last-child {
            color: var(--gray-800, #1F2A37);
            & svg {

            }
        }
    }
}
}
</style>