import React from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Home from './pages/Home'
import Profile from './pages/Profile'
import Layout from './components/Layout'
import { UserContextProvider } from './context/userContext'

const App: React.FC = () => {
  return (
    <UserContextProvider>
      <BrowserRouter>
        <div className='font-sans'>
          <Layout>
            <Routes>
              <Route path='/' element={<Home />} />
              <Route path='/profile' element={<Profile />} />
            </Routes>
          </Layout>
        </div>
      </BrowserRouter>
    </UserContextProvider>
  )
}

export default App
