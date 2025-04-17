import Calendar from "../components/Calendar"

const Home = () => {
  return (
    <div className="grid grid-flow-row-dense grid-cols-4 grid-rows-2 gap-5 mx-15 xl:mx-50 p-5 h-180">
      <section className="col-span-3 row-span-2 flex flex-col justify-center items-center shadow-xl p-10 rounded-xl">
        <h1 className="text-xl md:text-3xl font-medium">Hey, how was your day?</h1>
        <textarea className="border-2 border-gray-200 rounded-xl m-6 p-2 w-full max-w-2xl h-32 resize-none" placeholder="Share anything" />
      </section>

      <section className="shadow-xl p-5 rounded-xl">
        <Calendar/>
      </section>

      <section className="shadow-xl p-5 rounded-xl">
        Diaries
      </section>
    </div>
  )
}

export default Home