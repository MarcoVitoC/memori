<script setup lang="ts">
  import { ref } from 'vue'
  import Calendar from "../components/ui/Calendar.vue"
  import type { PostDiaryPayload } from "../types/diary"
  import { GetEnv } from "../utils/env"

  const apiUrl = GetEnv('VITE_API_URL', 'http://localhost:8000')

  const content = ref("")
  const setContent = (e: Event) => {
    content.value = (e.target as HTMLTextAreaElement).value
  }

  const handleKey = (e: KeyboardEvent) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault()
      postDiary()
    }
  }

  const handleSubmit = (e: Event) => {
    e.preventDefault()
    postDiary()
  }

  const postDiary = async () => {
    if (!content) return

    const payload: PostDiaryPayload = {
      content: content.value
    }

    try {
      const response = await fetch(`${apiUrl}/diaries`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
      })

      const data = await response.json()
      if (!response.ok) {
        throw new Error(data.errors)
      }
    } catch (error) {
      console.error('ERROR: ', error)
    }
  }
</script>

<template>
  <div class="grid grid-flow-row-dense grid-cols-4 grid-rows-2 gap-5 mx-5 xs:mx-15 xl:mx-40 p-5 sm:h-160">
    <section class="order-2 col-span-4 sm:order-3 lg:order-1 lg:col-span-3 lg:row-span-2 flex flex-col justify-center items-center shadow-xl p-7 sm:p-10 rounded-xl">
      <h1 class="text-xl sm:text-2xl lg:text-3xl font-medium text-center">Hey user, how was your day?</h1>
      <form class="flex flex-col items-center w-full" @submit="handleSubmit">
        <textarea class="border-2 border-gray-200 rounded-xl m-4 sm:m-6 p-4 w-full max-w-2xl h-32 resize-none" placeholder="Share anything" @keydown="handleKey" :value="content" @input="setContent" />
        <div class="flex justify-between w-full max-w-2xl p-4 -mt-22">
          <button class="border-2 border-gray-300 rounded-xl p-1 cursor-pointer hover:bg-gray-100">🖼️</button>
          <button type="submit" class="border-2 border-gray-300 rounded-xl p-1 cursor-pointer hover:bg-gray-100">⬆️</button>
        </div>
      </form>
    </section>

    <section class="order-1 col-span-4 sm:order-2 sm:col-span-2 lg:col-span-1 shadow-xl p-5 rounded-xl">
      <Calendar />
    </section>

    <section class="order-3 col-span-4 sm:order-1 sm:col-span-2 lg:order-3 lg:col-span-1 shadow-xl rounded-xl overflow-y-auto">
      <h1 class="border-b-2 border-gray-200 p-3 cursor-pointer hover:bg-slate-100">
        <span class="text-2xl">😭</span> Lorem ipsum dolor amet...
      </h1>
      <h1 class="border-b-2 border-gray-200 p-3 cursor-pointer hover:bg-slate-100">
        <span class="text-2xl">😥</span> Lorem ipsum dolor amet...
      </h1>
      <h1 class="border-b-2 border-gray-200 p-3 cursor-pointer hover:bg-slate-100">
        <span class="text-2xl">😱</span> Lorem ipsum dolor amet...
      </h1>
      <h1 class="border-b-2 border-gray-200 p-3 cursor-pointer hover:bg-slate-100">
        <span class="text-2xl">🤣</span> Lorem ipsum dolor amet...
      </h1>
      <h1 class="border-b-2 border-gray-200 p-3 cursor-pointer hover:bg-slate-100">
        <span class="text-2xl">🤬</span> Lorem ipsum dolor amet...
      </h1>
      <h1 class="border-b-2 border-gray-200 p-3 cursor-pointer hover:bg-slate-100">
        <span class="text-2xl">😐</span> Lorem ipsum dolor amet...
      </h1>
      <h1 class="border-b-2 border-gray-200 p-3 cursor-pointer hover:bg-slate-100">
        <span class="text-2xl">😴</span> Lorem ipsum dolor amet...
      </h1>
      <h1 class="border-b-2 border-gray-200 p-3 cursor-pointer hover:bg-slate-100">
        <span class="text-2xl">🤒</span> Lorem ipsum dolor amet...
      </h1>
    </section>
  </div>
</template>