<script setup lang="ts">
  import { DAYS, MONTHS_NAME, newCalendar } from "../../utils/calendar"

  const today = new Date()
  const month = today.getMonth()
  const year = today.getFullYear()
  
  const calendar = newCalendar(year, month)
</script>

<template>
  <div class="flex justify-between items-center mb-3">
    <h1 class="text-xl font-medium">{{ MONTHS_NAME[month] }} {{ year }}</h1>
    <h1 class="text-4xl">🙂</h1>
  </div>
  <div class="grid grid-cols-7 grid-rows-6 gap-3">
    <h1 
      v-for="(day, index) in DAYS"
      :key="index"
      class="grid place-content-center text-gray-400"
    >
      {{ day }}
    </h1>

    <h1 
      v-for="(day, index) in calendar"
      :key="index"
      :class="[
        day.isCurrentMonth ? '': 'text-gray-300',
        'grid place-content-center cursor-pointer'
      ]"
    >
      {{ day.date }}
    </h1>
  </div>
</template>