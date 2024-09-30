import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import DialogUI from '../components/UI/dialog';
import settings from '../assets/settings.svg';
import manage_account from '../assets/manage_account.svg';
import back from '../assets/back.svg';
import edit from '../assets/edit.svg';
import block from '../assets/block.svg';
import { ToolTip } from '../components/UI/tooltip';
import { userGeneral } from '../utils/userprops';
import { displaydate } from '../utils/general';

const Profile: React.FC = () => {
    const [user, setUser] = useState<userGeneral>();
    const navigate = useNavigate();
    const tabs = [
        { label: "General", logo: manage_account },
        { label: "Preferences", logo: settings }
    ];
    const [activeTab, setActiveTab] = useState<string>(tabs[0].label);


    const goBack = () => {
        navigate(-1);
    };

    useEffect(() => {
        setUser({
            username: "johndoe123",
            email: "johndoe@example.com",
            profile: {
                firstName: "John",
                lastName: "Doe",
                profilepicture: "https://static.vecteezy.com/system/resources/previews/004/607/791/non_2x/man-face-emotive-icon-smiling-male-character-in-blue-shirt-flat-illustration-isolated-on-white-happy-human-psychological-portrait-positive-emotions-user-avatar-for-app-web-design-vector.jpg"
            },
            dob: new Date("2010-05-05"),
            mobile: "1234567890",
            gender: "Male",
            preferences: {
                email: true,
                app: false,
                theme: true
            }
        })
    }, [])

    return (
        <div className='flex'>
            <div className='flex w-1/5 border-r min-h-screen'>
                <div className='flex flex-col w-full'>
                    <button
                        onClick={goBack}
                        className='text-lg font-bold p-2 text-black hover:text-gray-500'>
                        <img src={back} alt='back' className='w-7 h-7' />
                    </button>
                    <div className='flex text-3xl items-center justify-center mb-7'>SETTINGS</div>
                    {tabs.map((tab, index) => (
                        <div
                            key={index}
                            className={`flex items-center gap-x-2 py-3 pl-5 cursor-pointer ${activeTab === tab.label ? 'bg-gray-300' : 'hover:bg-gray-100'}`}
                            onClick={() => setActiveTab(tab.label)}
                        >
                            <img src={tab.logo} alt={`${tab.label} logo`} className='w-6 h-6' />
                            <button className='text-sm font-medium'>{tab.label}</button>
                        </div>
                    ))}
                </div>
            </div>

            <div className='flex w-4/5 p-5'>
                <div className='relative w-full'>
                    <div
                        className={`absolute w-full transition-opacity duration-500 ease-in-out ${activeTab === "General" ? "opacity-100" : "opacity-0 pointer-events-none"}`}
                        style={{ display: activeTab === "General" ? 'block' : 'none' }}
                    >
                        <div className='flex flex-col w-full'>
                            <div className='flex justify-center text-5xl font-extralight mt-6 mb-16'>Manage your account</div>
                            <div className='flex gap-x-4 items-start'>

                                <div className='w-1/2 flex flex-col rounded-md border p-4 gap-y-4'>
                                    <div className='text-xl'>Basic Information</div>
                                    <div className='font-light mb-2'>Some info may be visible to other people using College Book Community.</div>
                                    <div className='flex flex-col justify-center items-center'>
                                        <img src={user?.profile.profilepicture} alt='profile' className='w-20 h-20 mb-3' />
                                        <DialogUI
                                            title="Select image"
                                            inputLabel="Upload"
                                            description=""
                                            inputValue=""
                                            inputType="file"
                                        >
                                            <div className='px-1.5 py-0.5 mb-1 border rounded-md hover:bg-slate-50'>Edit</div>
                                        </DialogUI>
                                    </div>
                                    <div className='flex border-t justify-between items-center p-1'>
                                        <div className='font-light'>Userame</div>
                                        <div className='text-md'>{user?.username}</div>
                                        <DialogUI
                                            title="Change Username"
                                            inputLabel="Username"
                                            description=""
                                            inputValue=""
                                            inputType="text"
                                        >
                                            <img src={edit} alt='edit' className='w-4 h-4' />
                                        </DialogUI>
                                    </div>
                                    <div className='flex border-t justify-between items-center p-1'>
                                        <div className='font-light'>Birthday</div>
                                        <div className='text-md'>{displaydate(user?.dob)}</div>
                                        <DialogUI
                                            title="Change Birthday"
                                            inputLabel="Birthday"
                                            description=""
                                            inputValue=""
                                            inputType="date"
                                        >
                                            <img src={edit} alt='edit' className='w-4 h-4' />
                                        </DialogUI>
                                    </div>
                                    <div className='flex border-t justify-between items-center p-1'>
                                        <div className='font-light'>Gender</div>
                                        <div className='text-md'>{user?.gender}</div>
                                        <DialogUI
                                            title="Change Gender"
                                            inputLabel="Gender"
                                            description=""
                                            inputValue=""
                                            inputType="string"
                                        >
                                            <img src={edit} alt='edit' className='w-4 h-4' />
                                        </DialogUI>
                                    </div>
                                </div>

                                <div className='w-1/2 flex flex-col gap-y-4'>
                                    <div className='flex flex-col rounded-md border p-4 gap-y-4'>
                                        <div className='text-xl'>Contact Information</div>
                                        <div className='flex border-t justify-between items-center p-1'>
                                            <div className='font-light'>Email</div>
                                            <div className='text-md'>{user?.email}</div>
                                            <ToolTip message="You can't change your registered email "><img src={block} alt='edit' className='w-4 h-4' /></ToolTip>
                                        </div>
                                        <div className='flex border-t justify-between items-center p-1'>
                                            <div className='font-light'>Phone</div>
                                            <div className='text-md'>+91 {user?.mobile}</div>
                                            <DialogUI
                                                title="Change Mobile Number"
                                                inputLabel="Mobile Number"
                                                description=""
                                                inputValue=""
                                                inputType="string"
                                            >
                                                <img src={edit} alt='edit' className='w-4 h-4' />
                                            </DialogUI>
                                        </div>
                                    </div>

                                    <div className='flex flex-col rounded-md border p-4 gap-y-4'>
                                        <div className='text-xl'>Change Password</div>
                                        <div className='flex border-t justify-between items-center p-1'>
                                            <div className='font-light'>Password</div>
                                            <div className='text-md'>********</div>
                                            <DialogUI
                                                title="Change Password"
                                                inputLabel="Password"
                                                description=""
                                                inputValue=""
                                                inputType="password"
                                            >
                                                <img src={edit} alt='edit' className='w-4 h-4' />
                                            </DialogUI>
                                        </div>
                                    </div>

                                    <div className='flex flex-col rounded-md border p-4 gap-y-4'>
                                        <div className='text-xl text-red-600'>Delete My Account</div>
                                        <div className='flex border-t justify-between items-center p-1'>
                                            <div className='font-light'>Delete</div>
                                            <button className='px-2 py-1 bg-red-600 rounded-md text-white hover:bg-red-700'>
                                                Delete My Account
                                            </button>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div
                        className={`absolute w-full transition-opacity duration-500 ease-in-out ${activeTab === "Preferences" ? "opacity-100" : "opacity-0 pointer-events-none"}`}
                        style={{ display: activeTab === "Preferences" ? 'block' : 'none' }}
                    >
                        <div className='flex flex-col w-full'>
                            <div className='flex justify-center text-5xl font-extralight mt-6 mb-16'>Manage your Preferences</div>
                            <div className='flex flex-col gap-y-4'>
                                <div className='flex flex-col rounded-md border p-4'>
                                    <div className='text-xl'>Notifications</div>
                                    <div className='font-light mb-2'>Change your notification settings.</div>

                                </div>
                                <div className='flex flex-col rounded-md border p-4'>
                                    <div className='text-xl'>Theme</div>
                                    <div className='font-light mb-2'>Change your theme.</div>
                                </div>

                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Profile;
