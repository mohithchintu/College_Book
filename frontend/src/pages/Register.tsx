import React, { useState } from 'react'
import { FcGoogle } from "react-icons/fc";
import { Link } from 'react-router-dom';

const Register: React.FC = () => {
    const [mail, setMail] = useState<string>("");
    const [OTPcheck, setOTPcheck] = useState<boolean>(false);
    const sendOTP = () => {
        setOTPcheck(true)
    }
    return (
        <div className='flex justify-center items-center h-[34rem]'>
            {OTPcheck ? (
                <div className='flex flex-col gap-y-6 border rounded-lg shadow-md p-8'>
                    <div className='font-semibold text-3xl'>OTP Verification</div>
                    <div className='flex flex-col gap-y-2'>
                        <input type="text" placeholder="Email" className="rounded-lg border p-2" />
                        <button
                            className='flex justify-center gap-x-2 p-2 border rounded-lg items-center bg-blue-600 text-white
                    hover:bg-blue-700'
                        >
                            Verify
                        </button>
                    </div>
                </div>
            ) : (
                <div className='flex flex-col gap-y-6 border rounded-lg shadow-md p-8'>
                    <div className='font-semibold text-3xl'>Register for College Book</div>
                    <button className='flex justify-center gap-x-2 p-2 border rounded-lg items-center hover:bg-slate-50'>
                        <FcGoogle size={22} /> <>Continue with Google</>
                    </button>
                    <div className='flex justify-center'>
                        OR
                    </div>
                    <div className='flex flex-col gap-y-2'>
                        <input type="text" placeholder="Email" value={mail} onChange={(e) => setMail(e.target.value)} className="rounded-lg border p-2" />
                        <button
                            className='flex justify-center gap-x-2 p-2 border rounded-lg items-center bg-blue-600 text-white
                    hover:bg-blue-700'
                            onClick={sendOTP}
                        >
                            Register
                        </button>
                    </div>
                    <Link to='/login' className='flex font-extralight justify-center hover:underline'>Already have an account ?</Link>
                </div>
            )}

        </div>
    )
}

export default Register
