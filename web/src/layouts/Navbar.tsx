import { useState } from "react"
import { NavLink } from "react-router"
import cn from "../utils/cn"
import IconHamburgerMenu from "../assets/hamburger-menu.svg"
import IconX from "../assets/x.svg"

const Navbar = () => {
  const [isNavbarOpen, setIsNavbarOpen] = useState<boolean>(false)
  const handleNavbar = () => setIsNavbarOpen(!isNavbarOpen)

  return (
    <>
      <nav className="sticky top-0 border-b-2 border-gray-200 bg-white">
        <div className="flex justify-between items-center mx-5 xs:mx-15 xl:mx-40 p-5">
          <a href="" className="text-3xl font-bold font-quicksand">memori</a>
          <button onClick={handleNavbar} className="cursor-pointer">
            <img src={isNavbarOpen ? IconX : IconHamburgerMenu} alt="Menu" className="size-8" />
          </button>
        </div>

        {isNavbarOpen && <NavbarContent handleNavbar={handleNavbar}/>}
      </nav>
    </>
  )
}

const NavbarContent = ({ handleNavbar }: { handleNavbar: () => void }) => {
  return (
    <div className="fixed w-full h-full backdrop-blur-md">
      <div className="flex flex-col h-full justify-center items-center gap-10 pb-50 text-lg font-medium">
        <NavLink 
          to="/home"
          onClick={handleNavbar}
          className={({ isActive }) => cn(
            (isActive ? "" : "text-gray-300 hover:underline hover:underline-offset-4"),
            "text-4xl sm:text-5xl"
          )}
        >
          Home
        </NavLink>
        <NavLink 
          to="/diaries"
          onClick={handleNavbar}
          className={({ isActive }) => cn(
            (isActive ? "" : "text-gray-300 hover:underline hover:underline-offset-4"),
            "text-4xl sm:text-5xl"
          )}
        >
          Diaries
        </NavLink>
      </div>
    </div>
  )
}

export default Navbar