<template>
    <div class="sdf">
        <div class="gwgw flex" ref="FTW">
            <slot />
        </div>
    </div>
</template>
<script lang="ts" setup>

const emits = defineEmits(["close"])
const FTW = ref()
// useHead({
//     bodyAttrs: {
//         class: "locked",
//     }
// })
onClickOutside(FTW, (e) => {
    useHead({
        bodyAttrs: {
            class: "",
        }
    })
    emits('close')
})

</script>
<style lang="sass">
.sdf
    position: fixed
    top: 0
    display: grid
    place-items: center
    width: 100%
    // overflow-x: hidden
    // height: 100vh
    inset: 0
    background-color: rgba(#333, .9)
    .gwgw
        background-color: white
        transition: all .2s ease-in
        padding: min(5rem, 5vw)
        display: grid
        text-align: center
</style>