<template>
    <NuxtLayout>
        <NuxtPage />
    </NuxtLayout>
</template>
<script setup lang="ts">
useHead({
    title: "Shini Poker bro",
    meta: [{
        name: 'description',
        content: 'Отлично все',
    }, {
        name: 'keywords',
        content: 'Покерь, Poker'
    }]
})
</script>
<style lang="sass">
    .page-enter-active, .page-leave-active
        transition: all 0.2s

    .page-enter-from
        opacity: 0
        transform: translateX(-1rem)
        
    .page-leave-to
        opacity: 0
        transform: translateX(1rem)
</style>