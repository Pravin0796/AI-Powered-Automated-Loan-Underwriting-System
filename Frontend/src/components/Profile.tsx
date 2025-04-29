import { useEffect, useState } from "react";
import { userClient } from "../services/Grpc";
import { format } from "date-fns";
import { UserDetailsResponse } from "../proto/user";
import { getUserId } from "../utils/Auth";

const Profile = () => {
  const [userDetails, setUserDetails] = useState<UserDetailsResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const userId = parseInt(getUserId() || "0");

  useEffect(() => {
    const getUserDetails = async () => {
      if (!userId) {
        setError("User not logged in.");
        setLoading(false);
        return;
      }

      try {
        const response = await userClient.GetUserDetails({ userId });
        setUserDetails(response);
      } catch (err) {
        setError("Failed to fetch user details.");
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    getUserDetails();
  }, [userId]);

  if (loading) {
    return (
      <div className="flex justify-center items-center h-screen text-xl font-semibold">
        Loading your profile...
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex justify-center items-center h-screen text-red-500 text-lg">
        {error}
      </div>
    );
  }

  return (
    <div className="p-6 max-w-4xl mx-auto bg-gray-50 shadow-lg rounded-lg text-gray-800 mt-15">
      <h1 className="text-2xl font-bold mb-4">My Profile</h1>
      <div className="grid grid-cols-1 sm:grid-cols-2 gap-6">
        <div>
          <span className="font-semibold text-gray-600">Full Name:</span>
          <p>{userDetails?.fullName || "N/A"}</p>
        </div>
        <div>
          <span className="font-semibold text-gray-600">Email:</span>
          <p>{userDetails?.email || "N/A"}</p>
        </div>
        <div>
          <span className="font-semibold text-gray-600">Phone:</span>
          <p>{userDetails?.phone || "N/A"}</p>
        </div>
        <div>
          <span className="font-semibold text-gray-600">Address:</span>
          <p>{userDetails?.address || "N/A"}</p>
        </div>
        <div>
          <span className="font-semibold text-gray-600">Date of Birth:</span>
          <p>{userDetails?.dateOfBirth ? format(new Date(userDetails.dateOfBirth.seconds * 1000), "dd MMM yyyy") : "N/A"}</p>
        </div>
        <div>
          <span className="font-semibold text-gray-600">Credit Score:</span>
          <p>{userDetails?.creditScore || "N/A"}</p>
        </div>
      </div>
    </div>
  );
};

export default Profile;