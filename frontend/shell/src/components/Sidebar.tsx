import { NavLink } from 'react-router-dom';

function Sidebar() {
  return (
    <div className="bg-gray-800 text-white w-64 space-y-6 py-7 px-2 absolute inset-y-0 left-0 transform -translate-x-full md:relative md:translate-x-0 transition duration-200 ease-in-out">
      <div className="flex items-center space-x-2 px-4">
        <svg xmlns="http://www.w3.org/2000/svg" className="h-8 w-8" viewBox="0 0 20 20" fill="currentColor">
          <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-11a1 1 0 10-2 0v2H7a1 1 0 100 2h2v2a1 1 0 102 0v-2h2a1 1 0 100-2h-2V7z" clipRule="evenodd" />
        </svg>
        <span className="text-2xl font-extrabold">AI Agent</span>
      </div>
      
      <nav>
        <NavLink 
          to="/tasks" 
          className={({ isActive }) => 
            `block py-2.5 px-4 rounded transition duration-200 ${
              isActive ? 'bg-blue-500 text-white' : 'hover:bg-gray-700'
            }`
          }
        >
          Tasks
        </NavLink>
        <NavLink 
          to="/results" 
          className={({ isActive }) => 
            `block py-2.5 px-4 rounded transition duration-200 ${
              isActive ? 'bg-blue-500 text-white' : 'hover:bg-gray-700'
            }`
          }
        >
          Results
        </NavLink>
      </nav>
    </div>
  );
}

export default Sidebar;
