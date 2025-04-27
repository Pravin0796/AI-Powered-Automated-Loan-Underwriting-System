import React, { useEffect, useState } from "react";

const UserManagement = () => {
  interface User {
    id: number;
    name: string;
    email: string;
    role: string;
  }

  const [users, setUsers] = useState<User[]>([]);

  useEffect(() => {
    // Fetch users from API
    async function fetchUsers() {
      // Make API call
      setUsers([
        { id: 1, name: "John Doe", email: "john@example.com", role: "User" },
        { id: 2, name: "Jane Smith", email: "jane@example.com", role: "Admin" },
      ]);
    }
    fetchUsers();
  }, []);

  return (
    <div className="p-6">
      <h2 className="text-2xl font-bold mb-4">User Management</h2>
      <table className="w-full bg-white shadow-md rounded-lg">
        <thead className="bg-gray-100">
          <tr>
            <th className="p-2">ID</th>
            <th className="p-2">Name</th>
            <th className="p-2">Email</th>
            <th className="p-2">Role</th>
          </tr>
        </thead>
        <tbody>
          {users.map((user) => (
            <tr key={user.id} className="border-t text-center">
              <td className="p-2">{user.id}</td>
              <td className="p-2">{user.name}</td>
              <td className="p-2">{user.email}</td>
              <td className="p-2">{user.role}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default UserManagement;