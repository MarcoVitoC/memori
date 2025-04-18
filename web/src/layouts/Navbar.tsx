import { NavLink } from "react-router"
import cn from "../utils/cn"
import IconHamburgerMenu from "../assets/hamburger-menu.svg"
// import IconX from "../assets/x.svg"

const Navbar = () => {
  return (
    <>
      <nav className="border-b-2 border-gray-200">
        <div className="flex justify-between items-center mx-15 2xl:mx-50 p-5">
          <a href="" className="text-3xl font-bold font-quicksand">memori</a>
          <div className="flex items-center gap-10 text-lg font-medium">
            <NavLink 
              to="/home"
              className={({ isActive }) => cn(
                (isActive ? "bg-amber-200" : "bg-white"),
                "px-3 py-2 rounded-xl"
              )}
            >
              Home
            </NavLink>
            <NavLink 
              to="/diaries"
              className={({ isActive }) => cn(
                (isActive ? "bg-amber-200" : "bg-white"),
                "px-3 py-2 rounded-xl"
              )}
            >
              Diaries
            </NavLink>
          </div>
          <button className="cursor-pointer">
            <img src={IconHamburgerMenu} alt="Menu" className="size-8" />
          </button>
        </div>
      </nav>
    </>
  )
}

export default Navbar