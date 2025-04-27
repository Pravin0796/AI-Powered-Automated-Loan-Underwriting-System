import { NavLink } from "react-router-dom";

const Sidebar = () => {
  const links = [
    { name: "Dashboard", path: "/admin/dashboard" },
    { name: "Manage Loans", path: "/admin/loans" },
    { name: "User Management", path: "/admin/users" },
    { name: "Reports", path: "/admin/reports" },
    { name: "Notifications", path: "/admin/notifications" },
    { name: "Settings", path: "/admin/settings" },
  ];

  return (
    <aside className="w-64 bg-blue-600 text-white h-screen p-4">
      <h2 className="text-2xl font-bold mb-6">Admin Panel</h2>
      <nav>
        {links.map((link) => (
          <NavLink
            key={link.name}
            to={link.path}
            className={({ isActive }) =>
              `block py-2 px-4 rounded hover:bg-blue-700 ${isActive ? "bg-blue-800" : ""}`
            }
          >
            {link.name}
          </NavLink>
        ))}
      </nav>
    </aside>
  );
};

export default Sidebar;