import React, { useState } from 'react';

const Test: React.FC = () => {
    const [c1, setC1] = useState<boolean>(false);
    const [c2, setC2] = useState<boolean>(false);
    const [c3, setC3] = useState<boolean>(false);

    return (
        <div className='flex w-full p-10 space-x-2'>
            <div className='flex flex-col w-1/2 p-4 gap-y-2'>
                <div className='flex p-2 border justify-between'>
                    Box11
                    <button
                        className='p-1 rounded-full border'
                        onClick={() => setC1(!c1)}
                        aria-expanded={c1}
                    >
                        cli
                    </button>
                </div>
                {c1 && <div className='flex h-10 z-10'>content is here</div>}
                <div className='flex p-2 border justify-between'>
                    Box12
                    <button
                        className='p-1 rounded-full border'
                        onClick={() => setC2(!c2)}
                        aria-expanded={c2}
                    >
                        cli
                    </button>
                </div>
                {c2 && <div className='flex h-10 z-10'>content is here</div>}
                <div className='flex p-2 border'>Box13</div>
                <div className='flex p-2 border'>Box14</div>
            </div>
            <div className='flex flex-col w-1/2 p-4 gap-y-2'>
                <div className='flex p-2 border justify-between'>
                    Box21
                    <button
                        className='p-1 rounded-full border'
                        onClick={() => setC3(!c3)}
                        aria-expanded={c3}
                    >
                        cli
                    </button>
                </div>
                {c3 && <div className='flex h-10 z-10'>content is here</div>}
                <div className='flex p-2 border'>Box22</div>
                <div className='flex p-2 border'>Box23</div>
                <div className='flex p-2 border'>Box24</div>
            </div>
        </div>
    );
};

export default Test;
