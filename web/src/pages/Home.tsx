import Calendar from "../components/Calendar"

const Home = () => {
  return (
    <div className="grid grid-flow-row-dense grid-cols-4 grid-rows-2 gap-5 mx-5 xs:mx-15 xl:mx-40 p-5 sm:h-160">
      <section className="order-2 col-span-4 sm:order-3 lg:order-1 lg:col-span-3 lg:row-span-2 flex flex-col justify-center items-center shadow-xl p-7 sm:p-10 rounded-xl">
        <h1 className="text-xl sm:text-2xl lg:text-3xl font-medium text-center">Hey user, how was your day?</h1>
        <textarea className="border-2 border-gray-200 rounded-xl m-4 sm:m-6 p-2 w-full max-w-2xl h-32 resize-none" placeholder="Share anything" />
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