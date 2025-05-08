import { useState } from "react"
import Calendar from "../components/Calendar"
import { PostDiaryPayload } from "../types/diary"
import { GetEnv } from "../utils/env"

const Home = () => {
  const apiUrl = GetEnv('VITE_API_URL', 'http://localhost:8000')

  const [content, setContent] = useState("")

  const handleKey = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault()
      postDiary()
    }
  }

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    postDiary()
  }

  const postDiary = async () => {
    if (!content.trim()) return

    const payload: PostDiaryPayload = {
      content: content
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

  return (
    <div className="grid grid-flow-row-dense grid-cols-4 grid-rows-2 gap-5 mx-5 xs:mx-15 xl:mx-40 p-5 sm:h-160">
      <section className="order-2 col-span-4 sm:order-3 lg:order-1 lg:col-span-3 lg:row-span-2 flex flex-col justify-center items-center shadow-xl p-7 sm:p-10 rounded-xl">
        <h1 className="text-xl sm:text-2xl lg:text-3xl font-medium text-center">Hey user, how was your day?</h1>
        <form className="flex flex-col items-center w-full" onSubmit={handleSubmit}>
          <textarea className="border-2 border-gray-200 rounded-xl m-4 sm:m-6 p-4 w-full max-w-2xl h-32 resize-none" placeholder="Share anything" onKeyDown={handleKey} value={content} onChange={(e) => setContent(e.target.value)} />
          <div className="flex justify-between w-full max-w-2xl p-4 -mt-22">
            <button className="border-2 border-gray-300 rounded-xl p-1 cursor-pointer hover:bg-gray-100">🖼️</button>
            <button type="submit" className="border-2 border-gray-300 rounded-xl p-1 cursor-pointer hover:bg-gray-100">⬆️</button>
          </div>
        </form>
      </section>

      <section className="order-1 col-span-4 sm:order-2 sm:col-span-2 lg:col-span-1 shadow-xl p-5 rounded-xl">
        <Calendar/>
      </section>

      <section className="order-3 col-span-4 sm:order-1 sm:col-span-2 lg:order-3 lg:col-span-1 shadow-xl rounded-xl overflow-y-auto">
        <h1 className="border-b-2 border-gray-200 p-3 cursor-pointer hover:bg-slate-100">
          <span className="text-2xl">😭</span> Lorem ipsum dolor amet...
        </h1>
        <h1 className="border-b-2 border-gray-200 p-3 cursor-pointer hover:bg-slate-100">
          <span className="text-2xl">😥</span> Lorem ipsum dolor amet...
        </h1>
        <h1 className="border-b-2 border-gray-200 p-3 cursor-pointer hover:bg-slate-100">
          <span className="text-2xl">😱</span> Lorem ipsum dolor amet...
        </h1>
        <h1 className="border-b-2 border-gray-200 p-3 cursor-pointer hover:bg-slate-100">
          <span className="text-2xl">🤣</span> Lorem ipsum dolor amet...
        </h1>
        <h1 className="border-b-2 border-gray-200 p-3 cursor-pointer hover:bg-slate-100">
          <span className="text-2xl">🤬</span> Lorem ipsum dolor amet...
        </h1>
        <h1 className="border-b-2 border-gray-200 p-3 cursor-pointer hover:bg-slate-100">
          <span className="text-2xl">😐</span> Lorem ipsum dolor amet...
        </h1>
        <h1 className="border-b-2 border-gray-200 p-3 cursor-pointer hover:bg-slate-100">
          <span className="text-2xl">😴</span> Lorem ipsum dolor amet...
        </h1>
        <h1 className="border-b-2 border-gray-200 p-3 cursor-pointer hover:bg-slate-100">
          <span className="text-2xl">🤒</span> Lorem ipsum dolor amet...
        </h1>
      </section>
    </div>
  )
}

export default Home