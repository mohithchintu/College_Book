import React from 'react'
import { Link } from 'react-router-dom'
import profile from '../assets/profile.svg';
import { useUserContext } from '../context/userContext';

const Header: React.FC = () => {
    const navitems = [
        { name: 'Home', link: '/' },
        { name: 'Study/Work', link: '/material' },
        { name: 'Memories', link: '/memories' },
    ]
    const { isAuthenticated } = useUserContext();
    return (
        <div className='flex justify-between h-16 border-b items-center px-5 shadow-md'>
            <div className='text-xl'>Logo</div>
            <div className='flex space-x-5 items-center border rounded-lg justify-center px-3 py-1 text-lg shadow-sm'>
                {navitems.map((item, index) => (
                    <Link key={index} to={item.link}>
                        <div className='text-gray-600 hover:text-gray-900'>{item.name}</div>
                    </Link>
                ))}
            </div>
            {isAuthenticated ? (
                <div className='flex space-x-5 items-center border rounded-lg px-3 py-1 text-lg shadow-sm'>
                    <Link to='/profile' className='text-gray-600 hover:text-gray-900'>
                        Profile
                    </Link>
                    <img src={profile} alt='profile' className='w-8 h-8' />
                </div>
            ) : (<div className='flex space-x-5 items-center border rounded-lg px-3 py-1 text-lg shadow-sm'>
                <Link to='/login' className='text-gray-600 hover:text-gray-900'>
                    Login
                </Link>
                <Link to='/register' className='text-gray-600 hover:text-gray-900'>
                    Register
                </Link>
            </div>
            )}
        </div>
    )
}

export default Header
