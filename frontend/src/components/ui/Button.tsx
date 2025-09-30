interface ButtonProps {
  type?: "submit" | "button" | "reset";
  name: string;
  disabled?: boolean;
  loading?: boolean;
  textColor?: string;
  bgColor?: string;
  hoverBgColor?: string;
  onClick?: () => void;
}

export default function Button({
  type,
  name,
  disabled = false,
  loading = false,
  textColor = "text-slate-100",
  bgColor = "bg-slate-600",
  hoverBgColor = "bg-slate-700",
  onClick,
}: ButtonProps) {
  return (
    <button
      type={type}
      onClick={onClick}
      disabled={disabled || loading}
      className={`btn border-0 ${bgColor} hover:${hoverBgColor} transition ease-in-out duration-300 ${textColor} w-full my-5 focus:outline-none`}
    >
      {name}
    </button>
  );
}
