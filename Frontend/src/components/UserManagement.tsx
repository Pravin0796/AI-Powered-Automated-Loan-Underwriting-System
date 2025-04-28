import React, { useEffect, useState } from "react";
import { userClient } from "../services/Grpc";
import { Loader2 } from "lucide-react";
import { UserDetailsResponse } from "../proto/user";


interface User {
  id: number;
  fullName: string;
  email: string;
  phone: string;
  address: string;
  createdAt?: string; // Make sure this is a string after conversion
  creditScore: number;
  dateOfBirth?: string; // Ensure dateOfBirth is a string
}

const UserManagement = () => {
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function fetchUsers() {
      try {
        const response = await userClient.GetAllUsers({});
        console.log("Fetched users:", response);

        // Map the response and handle Timestamp conversion
        const mappedUsers: User[] = response.users.map((user: UserDetailsResponse) => ({
          id: user.userId,
          fullName: user.fullName,
          email: user.email,
          phone: user.phone,
          address: user.address,
          // Convert Timestamp to Date string for createdAt
          createdAt: user.createdAt ? new Date(user.createdAt.seconds * 1000).toLocaleString() : "N/A",
          creditScore: user.creditScore,
          // Convert Timestamp to Date string for dateOfBirth
          dateOfBirth: user.dateOfBirth
            ? new Date(user.dateOfBirth.seconds * 1000).toLocaleDateString()
            : "N/A",  // Ensure it's either a date string or "N/A" if not available
        }));

        setUsers(mappedUsers);
      } catch (err) {
        console.error("Error fetching users:", err);
      } finally {
        setLoading(false);
      }
    }
    fetchUsers();
  }, []);

  if (loading) {
    return (
      <div className="flex justify-center items-center h-64">
        <Loader2 className="animate-spin h-8 w-8 text-gray-500" />
      </div>
    );
  }

  return (
    <div className="p-6 bg-gray-50 min-h-screen">
      <h2 className="text-2xl font-bold mb-4 text-blue-600">User Management</h2>
      <div className="overflow-x-auto">
        <table className="w-full bg-white shadow-md rounded-lg">
          <thead className="bg-blue-100">
            <tr className="border-b text-center">
              {/* <th className="p-3 text-left">ID</th> */}
              <th className="p-3 ">Full Name</th>
              <th className="p-3 ">Email</th>
              <th className="p-3 ">Phone</th>
              <th className="p-3 ">Address</th>
              {/* <th className="p-3 ">Created At</th> */}
              <th className="p-3 ">Credit Score</th>
              <th className="p-3 ">Date of Birth</th>
            </tr>
          </thead>
          <tbody>
            {users.map((user) => (
              <tr key={user.id} className="border-b text-center hover:bg-blue-50">
                {/* <td className="p-3">{user.id}</td> */}
                <td className="p-3">{user.fullName}</td>
                <td className="p-3">{user.email}</td>
                <td className="p-3">{user.phone}</td>
                <td className="p-3">{user.address}</td>
                {/* <td className="p-3">{user.createdAt || "N/A"}</td> */}
                <td className="p-3">{user.creditScore}</td>
                <td className="p-3">{user.dateOfBirth || "N/A"}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default UserManagement;
