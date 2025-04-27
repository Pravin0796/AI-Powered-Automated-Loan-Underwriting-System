import { useState } from 'react';

const Settings = () => {
    const [settings, setSettings] = useState({
      notifications: true,
      darkMode: false,
    });
  
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
      const { name, checked } = e.target;
      setSettings({ ...settings, [name]: checked });
    };
  
    return (
      <div className="max-w-2xl mx-auto p-6 bg-white shadow-md rounded-lg">
        <h2 className="text-2xl font-bold mb-4">Settings</h2>
        <div className="space-y-4">
          <label className="flex items-center space-x-2">
            <input
              type="checkbox"
              name="notifications"
              checked={settings.notifications}
              onChange={handleChange}
            />
            <span>Enable Notifications</span>
          </label>
          <label className="flex items-center space-x-2">
            <input type="checkbox" name="darkMode" checked={settings.darkMode} onChange={handleChange} />
            <span>Enable Dark Mode</span>
          </label>
        </div>
      </div>
    );
  };
  
  export default Settings;