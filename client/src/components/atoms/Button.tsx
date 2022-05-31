import React from "react";

interface IButton {
  classes?: string;
  onClick: React.MouseEventHandler<HTMLButtonElement>;
  Icon?: any;
  label?: string | React.ReactNode;
}

const Button: React.FC<IButton> = ({ classes, onClick, Icon, label }) => {
  return (
    <>
      <div
        className={`flex justify-center bg-blue-500 text-gray-200 rounded-md ${classes}`}
      >
        <button
          className="flex items-center justify-center gap-2 w-full text-xl font-semibold p-2"
          onClick={onClick}
        >
          {Icon && Icon}
          {label && label}
        </button>
      </div>
    </>
  );
};

export default Button;
