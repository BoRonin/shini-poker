<template>
    <div class="button">
        <slot />
    </div>
</template>
<style lang="sass">
.button
    display: grid
    align-content: center
    padding:.3rem .8rem
    border: 1px solid #444
    border-radius: .3rem
    text-align: center
    cursor: pointer
    transition: all .2s ease-in-out
    margin: .2rem
    font-size: clamp(.9rem, calc(.8rem + 0.22vw), 1.25rem)
    &:hover
        color: white
        background-color: #444
.fit
    width: fit-content
    margin: 0 auto
.dark
    background-color: #444
    color: white
    &:hover
        background-color: #222
        
</style>