<script setup lang="ts">
  import { ref } from 'vue'
  import IconX from '../assets/x.svg'
  import IconHamburgerMenu from '../assets/hamburger-menu.svg'
  import NavbarContent from './NavbarContent.vue'

  const isNavbarOpen = ref(false)
  
  const handleNavbar = () => {
    isNavbarOpen.value = !isNavbarOpen.value
  }
</script>

<template>
  <nav class="sticky top-0 border-b-2 border-gray-200 bg-white">
    <div class="flex justify-between items-center mx-5 xs:mx-15 xl:mx-40 p-5">
      <a href="" class="text-3xl font-bold font-quicksand">memori</a>
      <button @click=handleNavbar class="cursor-pointer">
        <img :src="isNavbarOpen ? IconX : IconHamburgerMenu" alt="Menu" class="size-8" />
      </button>
    </div>

    <NavbarContent v-if="isNavbarOpen" :handleNavbar="handleNavbar"/>
  </nav>
</template>