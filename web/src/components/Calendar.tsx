import { DAYS, MONTHS_NAME, newCalendar } from "../utils/calendar"
import cn from "../utils/cn"

const Calendar = () => {
  const today = new Date()
  const month = today.getMonth()
  const year = today.getFullYear()
  
  const calendar = newCalendar(year, month)
  
  return (
    <>
      <div className="flex justify-between items-center mb-3">
        <h1 className="text-xl font-medium">{MONTHS_NAME[month]} {year}</h1>
        <h1 className="text-4xl">🙂</h1>
      </div>
      <div className="grid grid-cols-7 grid-rows-6 gap-3">
        {DAYS.map((day: string, index: number) => (
          <h1 key={index} className="grid place-content-center text-gray-400">{day}</h1>
        ))}

        {calendar.map(({isCurrentMonth, date}, index: number) => (
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