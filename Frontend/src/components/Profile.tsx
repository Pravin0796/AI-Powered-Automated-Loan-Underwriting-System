// import React, { useEffect, useState } from "react";
// import { userClient } from "../services/Grpc"; // Your gRPC client
// import { UserDetailsResponse } from "../proto/user"; // Ensure you have the correct generated types
// import { format } from "date-fns";

// const Profile: React.FC = () => {
//   // Get userId from localStorage (or cookies, depending on where you store it)
//   const userId = localStorage.getItem("userId"); // Assuming the userId is stored here after login

//   const [profile, setProfile] = useState<UserDetailsResponse | null>(null);
//   const [loading, setLoading] = useState(true);

//   useEffect(() => {
//     const fetchProfile = async () => {
//       if (!userId) {
//         console.error("User ID not found");
//         setLoading(false);
//         return;
//       }

//       try {
//         const res = await userClient.GetUserDetails({ userId: parseInt(userId) });
//         setProfile(res);
//       } catch (error) {
//         console.error("Error fetching profile:", error);
//       } finally {
//         setLoading(false);
//       }
//     };

//     fetchProfile();
//   }, [userId]);

//   if (loading) {
//     return (
//       <div className="flex justify-center items-center h-full text-xl">
//         Loading...
//       </div>
//     );
//   }

//   return (
//     <div className="p-4 sm:p-8 max-w-3xl mx-auto">
//       <h1 className="text-2xl font-bold mb-6">Your Profile</h1>

//       {profile ? (
//         <div className="grid gap-4 text-sm sm:text-base">
//           <div>
//             <span className="font-medium text-gray-700">Full Name:</span>{" "}
//             {profile.fullName}
//           </div>
//           <div>
//             <span className="font-medium text-gray-700">Email:</span>{" "}
//             {profile.email}
//           </div>
//           <div>
//             <span className="font-medium text-gray-700">Phone:</span>{" "}
//             {profile.phone}
//           </div>
//           <div>
//             <span className="font-medium text-gray-700">Address:</span>{" "}
//             {profile.address}
//           </div>
//           <div>
//             <span className="font-medium text-gray-700">Date of Birth:</span>{" "}
//             {profile.dateOfBirth
//               ? format(
//                   new Date(profile.dateOfBirth.seconds * 1000), // Converting Timestamp to Date
//                   "dd MMM yyyy"
//                 )
//               : "N/A"}
//           </div>
//           <div>
//             <span className="font-medium text-gray-700">Credit Score:</span>{" "}
//             {profile.creditScore}
//           </div>
//           <div>
//             <span className="font-medium text-gray-700">Account Created:</span>{" "}
//             {profile.createdAt
//               ? format(
//                   new Date(profile.createdAt.seconds * 1000), // Converting Timestamp to Date
//                   "dd MMM yyyy"
//                 )
//               : "N/A"}
//           </div>
//           <div>
//             <span className="font-medium text-gray-700">Last Updated:</span>{" "}
//             {profile.updatedAt
//               ? format(
//                   new Date(profile.updatedAt.seconds * 1000), // Converting Timestamp to Date
//                   "dd MMM yyyy"
//                 )
//               : "N/A"}
//           </div>
//         </div>
//       ) : (
//         <div className="text-red-500">Unable to load profile.</div>
//       )}
//     </div>
//   );
// };

// export default Profile;


// src/pages/Profile.tsx
import { useEffect, useState } from "react";
import { userClient } from "../services/Grpc"; // Assuming you have grpc setup for user details

const Profile = () => {
  const [userDetails, setUserDetails] = useState<any>(null);
  const userId = 1; // Example user ID (replace with dynamic data)

  useEffect(() => {
    const getUserDetails = async () => {
      try {
        const response = await userClient.GetUserDetails({ userId: userId });
        setUserDetails(response);
      } catch (error) {
        console.error("Error fetching user details", error);
      }
    };

    getUserDetails();
  }, [userId]);

  if (!userDetails) return <div>Loading...</div>;

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-semibold mb-4">Profile</h1>
      <div className="bg-white p-6 rounded-lg shadow-md">
        <p><strong>Name:</strong> {userDetails.fullName}</p>
        <p><strong>Email:</strong> {userDetails.email}</p>
        <p><strong>Phone:</strong> {userDetails.phone}</p>
        <p><strong>Address:</strong> {userDetails.address}</p>
        <p><strong>Credit Score:</strong> {userDetails.creditScore}</p>
      </div>
    </div>
  );
};

export default Profile;
