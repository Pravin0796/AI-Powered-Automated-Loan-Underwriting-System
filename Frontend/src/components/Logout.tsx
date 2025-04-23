// components/Sidebar.jsx (inside your menu or as a separate button)
import { useNavigate } from "react-router-dom";
import { removeToken } from "../utils/Auth.ts";

export default function Sidebar() {
  const navigate = useNavigate();

  const handleLogout = () => {
    removeToken();
    navigate("/login");
  };

  return (
    <button
      onClick={handleLogout}
      className="text-red-500 hover:underline px-4 py-2 w-full text-left"
    >
      Logout
    </button>
  );
}
