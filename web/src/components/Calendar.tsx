import cn from "../utils/cn"

const Calendar = () => {
  const days = ["S", "M", "T", "W", "T", "F", "S"]
  const months = ["January","February","March","April","May","June","July","August","September","October","November","December"]

  const date = new Date()
  const month = date.getMonth()
  const year = date.getFullYear()

  const firstDayOfMonth = new Date(year, month, 1)
  const firstDateOfMonth = firstDayOfMonth.getDate()
  const firstDayOfWeek = firstDayOfMonth.getDay()

  const lastDayPrevMonth = new Date(year, month, 0).getDate()

  const lastDayOfMonth = new Date(year, month + 1, 0)
  const lastDateOfMonth = lastDayOfMonth.getDate()
  const lastDayOfWeek = lastDayOfMonth.getDay()
  
  const dates = []
  for (let i=firstDayOfWeek-1; i>=0; i--) {
    dates.push({isCurrentMonth: false, date: lastDayPrevMonth - i})
  }

  for (let i=firstDateOfMonth; i<=lastDateOfMonth; i++) {
    dates.push({isCurrentMonth: true, date: i})
  }

  for (let i=1; i<=6-lastDayOfWeek; i++) {
    dates.push({isCurrentMonth: false, date: i})
  }
  
  return (
    <>
      <div className="flex justify-between items-center mb-3">
        <h1 className="text-xl font-medium">{months[month]} {year}</h1>
        <h1 className="text-4xl">🙂</h1>
      </div>
      <div className="grid grid-cols-7 grid-rows-6 gap-3">
        {days.map((day: string, index: number) => (
          <h1 key={index} className="grid place-content-center text-gray-400">{day}</h1>
        ))}

        {dates.map(({isCurrentMonth, date}, index: number) => (
          <h1 
            key={index} 
            className={cn(
              (isCurrentMonth ? "": "text-gray-300"),
              "grid place-content-center cursor-pointer"
            )}
          >{date}</h1>
        ))}
      </div>
    </>
  )
}

export default Calendar