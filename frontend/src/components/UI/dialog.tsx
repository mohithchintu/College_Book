import React from 'react';
import * as Dialog from '@radix-ui/react-dialog';

type DialogProps = {
  children: React.ReactNode;
  title: string;
  inputLabel: string;
  description: string;
  inputValue: any;
  inputType: any;
};

const DialogUI: React.FC<DialogProps> = ({ children, title, inputLabel, description, inputValue, inputType }) => (
  <Dialog.Root>
    <Dialog.Trigger asChild>
      <button className="focus:outline-none">
        {children}
      </button>
    </Dialog.Trigger>
    <Dialog.Portal>
      <Dialog.Overlay className="bg-black/50 fixed inset-0" />
      <Dialog.Content className="fixed top-1/2 left-1/2 w-[90vw] max-w-md bg-white p-6 rounded-lg shadow-lg transform -translate-x-1/2 -translate-y-1/2 focus:outline-none">
        <Dialog.Title className="text-lg font-semibold border-b pb-2 mb-4">
          {title}
        </Dialog.Title>
        <Dialog.Description className="mb-4 text-gray-600">
          {description}
        </Dialog.Description>
        <div className="mb-4">
          <label htmlFor="dialog-input" className="block text-sm mb-2">
            {inputLabel}
          </label>
          <input
            id="dialog-input"
            type={inputType}
            defaultValue={inputValue}
            className="w-full p-2 border rounded focus:outline-none focus:ring-2 focus:ring-indigo-500"
          />
        </div>
        <div className="flex justify-end space-x-3">
          <Dialog.Close asChild>
            <button className="px-1.5 py-0.5 bg-indigo-600 text-white rounded-md hover:bg-indigo-500 focus:outline-none">
              Save
            </button>
          </Dialog.Close>
          <Dialog.Close asChild>
            <button
              className="px-1.5 py-0.5 bg-white border rounded-md hover:bg-slate-50 focus:outline-none"
              aria-label="Close"
            >
              Cancel
            </button>
          </Dialog.Close>
        </div>
      </Dialog.Content>
    </Dialog.Portal>
  </Dialog.Root>
);

export default DialogUI;
