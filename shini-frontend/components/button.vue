<template>
    <div class="button">
        <slot />
    </div>
</template>
<style lang="sass">
.button
    padding:.3rem 1rem
    border: 1px solid #444
    border-radius: .3rem
    text-align: center
    cursor: pointer
    transition: all .2s ease-in-out
    margin: .2rem
    &:hover
        color: white
        background-color: #444
.finishGameButton
    width: fit-content
    margin: 0 auto
</style>