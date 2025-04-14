import { NavLink } from "react-router"

const Navbar = () => {
  const navLinkStyle = (isActive: boolean) => `${isActive ? "bg-amber-200" : "bg-white"} px-3 py-2 rounded-xl`

  return (
    <>
      <nav className="border-b-2 border-gray-200">
        <div className="flex justify-between items-center mx-30 p-5">
          <a href="" className="text-3xl font-bold font-quicksand">memori</a>
          <div className="flex items-center gap-10 text-lg font-medium">
            <NavLink 
              to="/home"
              className={({ isActive }) => navLinkStyle(isActive)}>
              Home
            </NavLink>
            <NavLink 
              to="/diaries"
              className={({ isActive }) => navLinkStyle(isActive)}>
              Diaries
            </NavLink>
          </div>
          <h1>avatar</h1>
        </div>
      </nav>
    </>
  )
}

export default Navbar