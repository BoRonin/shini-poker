<template>
  <div class="div" v-if="!!wait">
    <div class="container page mx-auto" v-if="user === null">
      <LoginSignIn />
    </div>
    <div class="container page mx-auto" v-else>
      <!-- <slot /> -->
      {{ user }}
    </div>
  </div>
</template>

<script setup lang="ts">
import {User} from '@/types/User'
const {user} = useAuth()
const wait = ref(false)

onMounted(()=>{
  useFetch<User>('http://localhost:8080/user',{
    method: 'GET',
    credentials: 'include',
  }).then((response) => {
    user.value = response.data.value
    wait.value = true
  })

})


</script>