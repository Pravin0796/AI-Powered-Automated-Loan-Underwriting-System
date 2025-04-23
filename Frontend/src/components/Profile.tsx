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
    <div className="p-4 sm:p-8 max-w-4xl mx-auto">
      <h1 className="text-3xl font-bold mb-6 text-center sm:text-left">My Profile</h1>
      <div className="bg-white shadow-lg rounded-2xl p-6 sm:p-8 grid grid-cols-1 sm:grid-cols-2 gap-6 text-gray-800 text-base">
        <div>
          <span className="font-semibold text-gray-600">Full Name:</span>
          <p>{userDetails?.fullName}</p>
        </div>
        <div>
          <span className="font-semibold text-gray-600">Email:</span>
          <p>{userDetails?.email}</p>
        </div>
        <div>
          <span className="font-semibold text-gray-600">Phone:</span>
          <p>{userDetails?.phone}</p>
        </div>
        <div>
          <span className="font-semibold text-gray-600">Address:</span>
          <p>{userDetails?.address}</p>
        </div>
        <div>
          <span className="font-semibold text-gray-600">Date of Birth:</span>
          <p>
            {userDetails?.dateOfBirth
              ? format(new Date(userDetails.dateOfBirth.seconds * 1000), "dd MMM yyyy")
              : "N/A"}
          </p>
        </div>
        <div>
          <span className="font-semibold text-gray-600">Credit Score:</span>
          <p>{userDetails?.creditScore ?? "N/A"}</p>
        </div>
        <div>
          <span className="font-semibold text-gray-600">Account Created:</span>
          <p>
            {userDetails?.createdAt
              ? format(new Date(userDetails.createdAt.seconds * 1000), "dd MMM yyyy")
              : "N/A"}
          </p>
        </div>
        <div>
          <span className="font-semibold text-gray-600">Last Updated:</span>
          <p>
            {userDetails?.updatedAt
              ? format(new Date(userDetails.updatedAt.seconds * 1000), "dd MMM yyyy")
              : "N/A"}
          </p>
        </div>
      </div>
    </div>
  );
};

export default Profile;
