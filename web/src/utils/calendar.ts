import type { Calendar } from "../types/calendar"

export const DAYS = ["S", "M", "T", "W", "T", "F", "S"]

export const MONTHS_NAME = [
  "January", "February", "March", "April", "May", "June",
  "July", "August", "September", "October", "November", "December"
]

export const newCalendar = (year: number, month: number): Calendar[] => {
  const firstDayOfMonth = new Date(year, month, 1)
  const firstDateOfMonth = firstDayOfMonth.getDate()
  const firstDayOfWeek = firstDayOfMonth.getDay()

  const lastDayPrevMonth = new Date(year, month, 0).getDate()

  const lastDayOfMonth = new Date(year, month + 1, 0)
  const lastDateOfMonth = lastDayOfMonth.getDate()
  const lastDayOfWeek = lastDayOfMonth.getDay()
  
  const calendar: Calendar[] = []
  for (let i=firstDayOfWeek-1; i>=0; i--) {
    calendar.push({isCurrentMonth: false, date: lastDayPrevMonth - i})
  }

  for (let i=firstDateOfMonth; i<=lastDateOfMonth; i++) {
    calendar.push({isCurrentMonth: true, date: i})
  }

  for (let i=1; i<=6-lastDayOfWeek; i++) {
    calendar.push({isCurrentMonth: false, date: i})
  }

  return calendar
}