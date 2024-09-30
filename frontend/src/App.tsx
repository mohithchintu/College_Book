import React from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Home from './pages/Home'
import Profile from './pages/Profile'
import Layout from './components/Layout'
import { UserContextProvider } from './context/userContext'
import Material from './pages/Material'
import Register from './pages/Register'
import Login from './pages/Login'
import Test from './pages/Test'

const App: React.FC = () => {
  return (
    <UserContextProvider>
      <BrowserRouter>
        <div className='font-sans'>
          <Layout>
            <Routes>
              <Route path='/' element={<Home />} />
              <Route path='/profile' element={<Profile />} />
              <Route path='/material' element={<Material />} />
              <Route path='/test' element={<Test />} />

              <Route path='/register' element={<Register />} />
              <Route path='/login' element={<Login />} />
            </Routes>
          </Layout>
        </div>
      </BrowserRouter>
    </UserContextProvider>
  )
}

export default App
