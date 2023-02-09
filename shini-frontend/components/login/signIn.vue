<template>
    <div class="flex min-h-full items-center text-center justify-center py-12 px-4 sm:px-6 lg:px-8">
        <div class="w-full max-w-md space-y-8">
            <div>
                <Logo class="logo mx-auto hy-auto flex h-20 w-20" />
                <h2 class="mt-6 text-center text-3xl font-bold tracking-tight text-gray-900">Заходи. Не ссы.</h2>
                <a href="#" class="font-medium text-gray-600 hover:text-gray-400">Регистрации нет, так что если ты не в
                    клубе - ты лох.</a>
            </div>
            <form class="mt-8 space-y-6" action="#" method="POST">
                <input type="hidden" name="remember" value="true" />
                <div class="-space-y-px rounded-md shadow-sm">
                    <div>
                        <label for="login" class="sr-only">Логин</label>
                        <input id="login" name="login" type="text" required
                            class="relative block w-full appearance-none rounded-none rounded-t-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-slate-700 focus:outline-none focus:ring-indigo-500 sm:text-sm"
                            placeholder="Логин" />
                    </div>
                    <div>
                        <label for="password" class="sr-only">Пароль</label>
                        <input id="password" name="password" type="password" autocomplete="current-password" required=""
                            class="relative block w-full appearance-none rounded-none rounded-b-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-slate-700 focus:outline-none focus:ring-indigo-500 sm:text-sm"
                            placeholder="Пароль" />
                    </div>
                </div>

                <div>
                    <button type="submit"
                        class="group relative flex w-full justify-center rounded-md border border-transparent bg-slate-700 py-2 px-4 text-sm font-medium text-white hover:border-slate-900 hover:text-slate-900 hover:bg-slate-300 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
                        <!-- <span class="absolute inset-y-0 left-0 flex items-center pl-3">
                        </span> -->
                        В ЗАРУБУ!
                    </button>
                </div>
            </form>
        </div>
    </div>
</template>

<script setup>
</script>

<style lang="sass" scoped>
    button
        transition: all .3s ease-in-out
</style>