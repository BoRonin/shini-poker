<template>
  <div>
    <NuxtLayout>
      <transition name="page">

        <NuxtPage />
      </transition>


    </NuxtLayout>
  </div>
</template>

<style lang="sass">
.page-enter-active, .page-leave-active
    transition: all 0.2s

.page-enter-from
    opacity: 0
    transform: translateX(-1rem)
     
.page-leave-to
    opacity: 0
    transform: translateX(1rem)
</style>