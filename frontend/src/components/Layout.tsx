import React, { ReactNode } from 'react';
import { useLocation } from 'react-router-dom';
import Header from './Header';

const Layout: React.FC<{ children: ReactNode }> = ({ children }) => {
    const location = useLocation();
    const isProfilePage = location.pathname === '/profile';

    return (
        <div>
            {isProfilePage ? (
                <div>
                    <main>{children}</main>
                </div>
            ) : (
                <div>
                    <Header />
                    <main>{children}</main>
                </div>
            )}
        </div>
    );
};

export default Layout;
