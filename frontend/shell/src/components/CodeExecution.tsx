import React from 'react';

function CodeExecution() {
  return (
    <div className="bg-white shadow-md rounded-lg p-6">
      <h1 className="text-2xl font-bold text-gray-800 mb-4">Code Execution</h1>
      
      <div className="mb-6 p-4 bg-blue-50 border border-blue-200 rounded-md">
        <p className="text-blue-700">
          This is a placeholder for the Code Execution service. 
          In a production environment, this would allow you to execute code in various languages.
        </p>
      </div>
      
      <div className="space-y-4">
        <div>
          <label htmlFor="language" className="block text-sm font-medium text-gray-700 mb-1">
            Language
          </label>
          <select
            id="language"
            className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="python">Python</option>
            <option value="javascript">JavaScript</option>
            <option value="go">Go</option>
            <option value="bash">Bash</option>
          </select>
        </div>
        
        <div>
          <label htmlFor="code" className="block text-sm font-medium text-gray-700 mb-1">
            Code
          </label>
          <textarea
            id="code"
            rows={10}
            className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 font-mono"
            placeholder="Enter your code here..."
            defaultValue="# Example Python code
def hello_world():
    print('Hello, world!')
    
hello_world()"
          ></textarea>
        </div>
        
        <div className="flex justify-end">
          <button
            type="button"
            className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Execute Code
          </button>
        </div>
        
        <div>
          <label htmlFor="output" className="block text-sm font-medium text-gray-700 mb-1">
            Output
          </label>
          <div
            id="output"
            className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-gray-50 font-mono h-32 overflow-auto"
          >
            <p className="text-gray-500">Output will appear here after execution...</p>
          </div>
        </div>
      </div>
    </div>
  );
}

export default CodeExecution;
