import Calendar from "../components/Calendar"

const Home = () => {
  return (
    <div className="grid grid-flow-row-dense grid-cols-4 grid-rows-2 gap-5 mx-5 xs:mx-15 2xl:mx-50 p-5 sm:h-160">
      <section className="order-2 col-span-4 sm:order-3 lg:order-1 lg:col-span-3 lg:row-span-2 flex flex-col justify-center items-center shadow-xl p-7 sm:p-10 rounded-xl">
        <h1 className="text-lg sm:text-2xl md:text-3xl font-medium text-center">Hey user, how was your day?</h1>
        <textarea className="border-2 border-gray-200 rounded-xl m-4 sm:m-6 p-2 w-full max-w-2xl h-32 resize-none" placeholder="Share anything" />
      </section>

      <section className="order-1 col-span-4 sm:order-2 sm:col-span-2 lg:col-span-1 shadow-xl p-5 rounded-xl">
        <Calendar/>
      </section>

      <section className="order-3 col-span-4 sm:order-1 sm:col-span-2 lg:order-3 lg:col-span-1 shadow-xl p-5 rounded-xl overflow-y-auto">
        Diaries
        Lorem ipsum dolor sit amet consectetur adipisicing elit. Vitae corrupti rem harum fugit praesentium hic a! Maxime, vel natus veritatis illum, sunt fugit, repellat debitis iusto quos ipsa placeat hic. Lorem ipsum dolor sit amet, consectetur adipisicing elit. Temporibus, excepturi laboriosam sequi, possimus id, doloribus error officia quas non facilis cumque! Cupiditate sunt quas quasi harum repellat facilis deleniti ipsum?
      </section>
    </div>
  )
}

export default Home