// components/Sidebar.jsx (inside your menu or as a separate button)
import { useNavigate } from "react-router-dom";

export default function Sidebar() {
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("user_id");
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
