// import React, { useEffect, useState } from "react";

// const UserManagement = () => {
//   interface User {
//     id: number;
//     name: string;
//     email: string;
//     role: string;
//   }

//   const [users, setUsers] = useState<User[]>([]);

//   useEffect(() => {
//     // Fetch users from API
//     async function fetchUsers() {
//       // Make API call
//       setUsers([
//         { id: 1, name: "John Doe", email: "john@example.com", role: "User" },
//         { id: 2, name: "Jane Smith", email: "jane@example.com", role: "Admin" },
//       ]);
//     }
//     fetchUsers();
//   }, []);

//   return (
//     <div className="p-6">
//       <h2 className="text-2xl font-bold mb-4">User Management</h2>
//       <table className="w-full bg-white shadow-md rounded-lg">
//         <thead className="bg-gray-100">
//           <tr>
//             <th className="p-2">ID</th>
//             <th className="p-2">Name</th>
//             <th className="p-2">Email</th>
//             <th className="p-2">Role</th>
//           </tr>
//         </thead>
//         <tbody>
//           {users.map((user) => (
//             <tr key={user.id} className="border-t text-center">
//               <td className="p-2">{user.id}</td>
//               <td className="p-2">{user.name}</td>
//               <td className="p-2">{user.email}</td>
//               <td className="p-2">{user.role}</td>
//             </tr>
//           ))}
//         </tbody>
//       </table>
//     </div>
//   );
// };

// export default UserManagement;


import React, { useEffect, useState } from "react";

interface User {
  id: number;
  name: string;
  email: string;
  role: string;
}

const UserManagement = () => {
  const [users, setUsers] = useState<User[]>([]);

  useEffect(() => {
    // Simulated API call
    setUsers([
      { id: 1, name: "John Doe", email: "john@example.com", role: "User" },
      { id: 2, name: "Jane Smith", email: "jane@example.com", role: "Admin" },
    ]);
  }, []);

  return (
    <div className="p-6 bg-gray-50 min-h-screen">
      <h2 className="text-2xl font-bold mb-4 text-blue-600">User Management</h2>
      <div className="overflow-x-auto">
        <table className="w-full bg-white shadow-md rounded-lg">
          <thead className="bg-blue-100">
            <tr>
              <th className="p-3 text-left">ID</th>
              <th className="p-3 text-left">Name</th>
              <th className="p-3 text-left">Email</th>
              <th className="p-3 text-left">Role</th>
            </tr>
          </thead>
          <tbody>
            {users.map((user) => (
              <tr key={user.id} className="border-t hover:bg-blue-50">
                <td className="p-3">{user.id}</td>
                <td className="p-3">{user.name}</td>
                <td className="p-3">{user.email}</td>
                <td className="p-3">{user.role}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default UserManagement;