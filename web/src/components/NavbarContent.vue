<script setup lang="ts">
import { RouterLink, useRoute } from 'vue-router';

const route = useRoute()

const { handleNavbar } = defineProps<{
  handleNavbar: () => void
}>()
</script>

<template>
  <div class="fixed w-full h-full backdrop-blur-md">
    <div class="flex flex-col h-full justify-center items-center gap-10 pb-50 text-lg font-medium">
      <RouterLink 
        to="/home" 
        @click="handleNavbar"
        :class="[
          route.path === '/home' ? 'text-black' : 'text-gray-300',
          'hover:text-black text-4xl sm:text-5xl'
        ]"
      >
        Home
      </RouterLink>
      
      <RouterLink 
        to="/diaries" 
        @click="handleNavbar" 
        :class="[
          route.path === '/diaries' ? 'text-black' : 'text-gray-300',
          'hover:text-black text-4xl sm:text-5xl'
        ]"
      >
        Diaries
      </RouterLink>
    </div>
  </div>
</template>