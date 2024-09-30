import * as Tooltip from '@radix-ui/react-tooltip';
import { ReactNode } from 'react';

export const ToolTip: React.FC<{ children: ReactNode, message: string }> = ({ children, message }) => (
    <Tooltip.Provider>
        <Tooltip.Root>
            <Tooltip.Trigger>
                {children}
            </Tooltip.Trigger>
            <Tooltip.Portal>
                <Tooltip.Content
                    sideOffset={5}
                    className='data-[state=delayed-open]:data-[side=top]:animate-slideDownAndFade data-[state=delayed-open]:data-[side=right]:animate-slideLeftAndFade data-[state=delayed-open]:data-[side=left]:animate-slideRightAndFade data-[state=delayed-open]:data-[side=bottom]:animate-slideUpAndFade text-violet11 select-none rounded-[4px] bg-white px-[15px] py-[10px] text-[15px] leading-none shadow-[hsl(206_22%_7%_/_35%)_0px_10px_38px_-10px,_hsl(206_22%_7%_/_20%)_0px_10px_20px_-15px] will-change-[transform,opacity]'>
                    {message}
                    <Tooltip.Arrow className='fill-white' />
                </Tooltip.Content>
            </Tooltip.Portal>
        </Tooltip.Root>
    </Tooltip.Provider>
);
