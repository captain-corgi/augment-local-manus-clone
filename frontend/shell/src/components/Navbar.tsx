function Navbar() {
  return (
    <header className="bg-white shadow">
      <div className="mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16">
          <div className="flex">
            <div className="flex-shrink-0 flex items-center">
              <h1 className="text-xl font-bold text-gray-900">Local AI Agent System</h1>
            </div>
          </div>
          <div className="flex items-center">
            <span className="inline-flex items-center px-3 py-0.5 rounded-full text-sm font-medium bg-green-100 text-green-800">
              Running
            </span>
          </div>
        </div>
      </div>
    </header>
  );
}

export default Navbar;
