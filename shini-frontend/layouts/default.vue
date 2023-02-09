<template>
  <div class="container page mx-auto" v-if="user === null">
    <LoginSignIn />
  </div>
  <div class="container page mx-auto" v-else>
    <!-- <slot /> -->
    {{ user }}
  </div>
</template>

<script setup lang="ts">
const {user} = useAuth()

</script>