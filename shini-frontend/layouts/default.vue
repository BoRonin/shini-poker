<template>
    <div class="divlayout" v-if="!!wait">
        <transition name="element" mode="out-in">
            <div class="page mx-auto" v-if="user === null">
                <LoginSignIn />
            </div>
            <div class="container page mx-auto" v-else>
                <Header />
                <slot />
            </div>
        </transition>
</div>
</template>

<script setup lang="ts">
import { User } from '@/types/User'
const config = useAppConfig()
const { user } = useAuth()
const wait = ref(false)
onMounted(() => {
    useFetch<User>(config.BASE_URL + 'user', {
        method: 'GET',
        credentials: 'include',
    }).then((response) => {
        user.value = response.data.value
        wait.value = true
    })

})


</script>

<style lang="sass" >
// .divlayout 
//     background-color: #f4f4f4

</style>