import { BrowserRouter, Routes, Route } from 'react-router'
import Navbar from "./layouts/Navbar"
import Home from './pages/Home'
import Diaries from './pages/Diaries'
// import Footer from "./layouts/Footer"

function App() {

  return (
    <>
      <BrowserRouter>
        <Navbar/>
        <Routes>
          <Route path='/home' element={<Home />}/>
          <Route path='/diaries' element={<Diaries />}/>
        </Routes>
        {/* <Footer/> */}
      </BrowserRouter>
    </>
  )
}

export default App
