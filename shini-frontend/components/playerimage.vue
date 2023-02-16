<template>
    <img :src="source" :alt="alt" :class="class">
</template>
<script setup lang="ts">
    const props = defineProps({
        source: String,
        alt: String,
        class: String,
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
        user-select: none
    .small
        width: 5em
</style>