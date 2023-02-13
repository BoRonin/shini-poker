<template>
    <div class="container header flex mx-auto">
            <Logo class="logo hy-auto flex h-20 w-20" />
        <div class="user-header flex items-center gap-3">
            <p class="user" >{{ user?.login }}</p>
            <button @click="Logout()" class="logout_button">НА ВЫХОД</button>
        </div>
    </div>
    <hr/>
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
</style>