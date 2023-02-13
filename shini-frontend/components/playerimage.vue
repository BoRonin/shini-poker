<template>
    <img :src="source" :alt="alt">
</template>
<script setup lang="ts">
    const props = defineProps({
        source: String,
        alt: String
    })
</script>

<style lang="sass" scoped>
    img
        width: 10em
        aspect-ratio: 1/1
        object-fit: cover
        object-position: top
        border-radius: 50%
        margin: 0 auto
</style>