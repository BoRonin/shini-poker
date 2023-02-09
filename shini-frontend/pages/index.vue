<template>
    <div class="conatiner mx-auto text-center">
        <div class="login-section hy-auto mt-40">
            <LoginSignIn />
        </div>
    </div>
</template>

<script setup lang="ts">
const route = useRouter()
const config = useRuntimeConfig()
const stringurl = ref<string | null>()
stringurl.value = config.BASE_URL


</script>

<style lang="sass" scoped>

</style>