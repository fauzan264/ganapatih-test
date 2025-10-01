"use client";
import useAuthStore from "@/store/useAuthStore";
import { useRouter } from "next/navigation";
import { FaUserCircle } from "react-icons/fa";

export default function Navbar() {
  const { logout } = useAuthStore();
  const auth = useAuthStore();
  const router = useRouter();

  const onLogout = () => {
    logout();
    router.push("/");
  };

  return (
    <div className="navbar fixed font-bold shadow-sm transition duration-300 left-0 top-0 z-99 px-10 bg-slate-900 text-gray-200">
      <div className="navbar-start gap-5">My Feed</div>
      <div className="navbar-end hidden lg:flex items-center gap-6">
        {auth.username && (
          <div className="dropdown dropdown-end">
            <div
              tabIndex={0}
              role="button"
              className="btn btn-ghost btn-md avatar flex"
            >
              <div className="w-7 rounded-full">
                <FaUserCircle className="w-full h-full" />
              </div>
              <span className="ml-1 my-auto">{auth.username}</span>
            </div>
            <ul
              tabIndex={0}
              className="menu menu-sm dropdown-content bg-slate-700 rounded-box z-1 mt-3 w-52 p-2 shadow text-slate-200"
            >
              <li>
                <button onClick={() => onLogout()}>Logout</button>
              </li>
            </ul>
          </div>
        )}
      </div>
    </div>
  );
}
