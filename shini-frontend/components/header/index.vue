<template>
    <div class="header_wrapper">
        <div class="container header flex mx-auto">
                <Logo class="logo hy-auto flex h-20 w-20" />
            <div class="user-header flex items-center gap-3">
                <p class="user" >{{ user?.login }}</p>
                <Button @click="Logout()" class=" dark">НА ВЫХОД</Button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">

const { user } = useAuth()
const config = useAppConfig()
const Logout = async () => {
    await useFetch(config.BASE_URL + "logout", {
        method: 'POST',
        credentials: 'include',
        onResponse({ request, response }) {
            user.value = null
            navigateTo("/")
        }
    })
}
</script>

<style lang="sass" scoped>
    .header
        align-items: center
        justify-content: space-between
        padding: 0 1rem
        border: 1px solid #e9e9e9
        .logo
            margin: 1rem 0
        .logout_button
            padding: .5rem 1rem
            border: 1px solid #888
            border-radius: .4rem
            background: white
            transition: all .2s ease-in-out
            &:hover
                border-color: #9909ba
                color: white
                background: #9909ba
.header_wrapper
    background-color: white
</style>