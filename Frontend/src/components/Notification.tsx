const Notifications = ({ notifications }: { notifications: string[] }) => {
    return (
      <div className="max-w-md mx-auto p-4 bg-white shadow-md rounded-lg">
        <h2 className="text-xl font-bold mb-4">Notifications</h2>
        {notifications.length > 0 ? (
          <ul className="space-y-2">
            {notifications.map((note, index) => (
              <li key={index} className="p-2 bg-gray-100 rounded">
                {note}
              </li>
            ))}
          </ul>
        ) : (
          <p className="text-gray-700">No new notifications.</p>
        )}
      </div>
    );
  };
  
  export default Notifications;