import React, { useState } from 'react';

function Settings() {
  const [activeTab, setActiveTab] = useState('general');
  
  return (
    <div className="bg-white shadow-md rounded-lg p-6">
      <h1 className="text-2xl font-bold text-gray-800 mb-6">Settings</h1>
      
      <div className="border-b border-gray-200">
        <nav className="-mb-px flex space-x-8">
          <button
            className={`${
              activeTab === 'general'
                ? 'border-blue-500 text-blue-600'
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            } whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm`}
            onClick={() => setActiveTab('general')}
          >
            General
          </button>
          <button
            className={`${
              activeTab === 'api'
                ? 'border-blue-500 text-blue-600'
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            } whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm`}
            onClick={() => setActiveTab('api')}
          >
            API Keys
          </button>
          <button
            className={`${
              activeTab === 'appearance'
                ? 'border-blue-500 text-blue-600'
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            } whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm`}
            onClick={() => setActiveTab('appearance')}
          >
            Appearance
          </button>
          <button
            className={`${
              activeTab === 'advanced'
                ? 'border-blue-500 text-blue-600'
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            } whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm`}
            onClick={() => setActiveTab('advanced')}
          >
            Advanced
          </button>
        </nav>
      </div>
      
      <div className="mt-6">
        {activeTab === 'general' && (
          <div className="space-y-6">
            <div>
              <h3 className="text-lg font-medium text-gray-900">System Information</h3>
              <p className="mt-1 text-sm text-gray-500">
                Basic information about your Local AI Agent System.
              </p>
              <div className="mt-4 border border-gray-200 rounded-md overflow-hidden">
                <div className="px-4 py-5 sm:p-6">
                  <dl className="grid grid-cols-1 gap-x-4 gap-y-8 sm:grid-cols-2">
                    <div className="sm:col-span-1">
                      <dt className="text-sm font-medium text-gray-500">Version</dt>
                      <dd className="mt-1 text-sm text-gray-900">1.0.0</dd>
                    </div>
                    <div className="sm:col-span-1">
                      <dt className="text-sm font-medium text-gray-500">Last Updated</dt>
                      <dd className="mt-1 text-sm text-gray-900">June 15, 2023</dd>
                    </div>
                    <div className="sm:col-span-1">
                      <dt className="text-sm font-medium text-gray-500">Environment</dt>
                      <dd className="mt-1 text-sm text-gray-900">Development</dd>
                    </div>
                    <div className="sm:col-span-1">
                      <dt className="text-sm font-medium text-gray-500">Status</dt>
                      <dd className="mt-1 text-sm text-gray-900">
                        <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                          Active
                        </span>
                      </dd>
                    </div>
                  </dl>
                </div>
              </div>
            </div>
            
            <div>
              <h3 className="text-lg font-medium text-gray-900">Notifications</h3>
              <p className="mt-1 text-sm text-gray-500">
                Manage how you receive notifications.
              </p>
              <div className="mt-4">
                <div className="space-y-4">
                  <div className="flex items-start">
                    <div className="flex items-center h-5">
                      <input
                        id="email-notifications"
                        name="email-notifications"
                        type="checkbox"
                        className="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded"
                        defaultChecked
                      />
                    </div>
                    <div className="ml-3 text-sm">
                      <label htmlFor="email-notifications" className="font-medium text-gray-700">
                        Email notifications
                      </label>
                      <p className="text-gray-500">Get notified when a task is completed.</p>
                    </div>
                  </div>
                  <div className="flex items-start">
                    <div className="flex items-center h-5">
                      <input
                        id="browser-notifications"
                        name="browser-notifications"
                        type="checkbox"
                        className="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded"
                        defaultChecked
                      />
                    </div>
                    <div className="ml-3 text-sm">
                      <label htmlFor="browser-notifications" className="font-medium text-gray-700">
                        Browser notifications
                      </label>
                      <p className="text-gray-500">Get browser notifications for important events.</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        )}
        
        {activeTab === 'api' && (
          <div className="space-y-6">
            <div>
              <h3 className="text-lg font-medium text-gray-900">API Keys</h3>
              <p className="mt-1 text-sm text-gray-500">
                Manage your API keys for external services.
              </p>
              <div className="mt-4 space-y-4">
                <div>
                  <label htmlFor="openai-api-key" className="block text-sm font-medium text-gray-700">
                    OpenAI API Key
                  </label>
                  <div className="mt-1 flex rounded-md shadow-sm">
                    <input
                      type="password"
                      name="openai-api-key"
                      id="openai-api-key"
                      className="flex-1 min-w-0 block w-full px-3 py-2 rounded-none rounded-l-md focus:ring-blue-500 focus:border-blue-500 sm:text-sm border-gray-300"
                      placeholder="sk-..."
                      defaultValue="sk-••••••••••••••••••••••••••••••••••••••••••••••••"
                    />
                    <button
                      type="button"
                      className="inline-flex items-center px-3 py-2 border border-l-0 border-gray-300 rounded-r-md bg-gray-50 text-gray-500 sm:text-sm"
                    >
                      Show
                    </button>
                  </div>
                </div>
                
                <div>
                  <label htmlFor="google-api-key" className="block text-sm font-medium text-gray-700">
                    Google API Key
                  </label>
                  <div className="mt-1 flex rounded-md shadow-sm">
                    <input
                      type="password"
                      name="google-api-key"
                      id="google-api-key"
                      className="flex-1 min-w-0 block w-full px-3 py-2 rounded-none rounded-l-md focus:ring-blue-500 focus:border-blue-500 sm:text-sm border-gray-300"
                      placeholder="AIza..."
                      defaultValue="AIza••••••••••••••••••••••••••••••••••••••••••••"
                    />
                    <button
                      type="button"
                      className="inline-flex items-center px-3 py-2 border border-l-0 border-gray-300 rounded-r-md bg-gray-50 text-gray-500 sm:text-sm"
                    >
                      Show
                    </button>
                  </div>
                </div>
                
                <div>
                  <label htmlFor="github-api-key" className="block text-sm font-medium text-gray-700">
                    GitHub API Key
                  </label>
                  <div className="mt-1 flex rounded-md shadow-sm">
                    <input
                      type="password"
                      name="github-api-key"
                      id="github-api-key"
                      className="flex-1 min-w-0 block w-full px-3 py-2 rounded-none rounded-l-md focus:ring-blue-500 focus:border-blue-500 sm:text-sm border-gray-300"
                      placeholder="ghp_..."
                    />
                    <button
                      type="button"
                      className="inline-flex items-center px-3 py-2 border border-l-0 border-gray-300 rounded-r-md bg-gray-50 text-gray-500 sm:text-sm"
                    >
                      Show
                    </button>
                  </div>
                </div>
              </div>
            </div>
            
            <div className="flex justify-end">
              <button
                type="button"
                className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              >
                Save API Keys
              </button>
            </div>
          </div>
        )}
        
        {activeTab === 'appearance' && (
          <div className="space-y-6">
            <div>
              <h3 className="text-lg font-medium text-gray-900">Theme</h3>
              <p className="mt-1 text-sm text-gray-500">
                Customize the appearance of your Local AI Agent System.
              </p>
              <div className="mt-4">
                <div className="space-y-4">
                  <div className="flex items-center">
                    <input
                      id="theme-light"
                      name="theme"
                      type="radio"
                      className="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300"
                      defaultChecked
                    />
                    <label htmlFor="theme-light" className="ml-3 block text-sm font-medium text-gray-700">
                      Light
                    </label>
                  </div>
                  <div className="flex items-center">
                    <input
                      id="theme-dark"
                      name="theme"
                      type="radio"
                      className="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300"
                    />
                    <label htmlFor="theme-dark" className="ml-3 block text-sm font-medium text-gray-700">
                      Dark
                    </label>
                  </div>
                  <div className="flex items-center">
                    <input
                      id="theme-system"
                      name="theme"
                      type="radio"
                      className="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300"
                    />
                    <label htmlFor="theme-system" className="ml-3 block text-sm font-medium text-gray-700">
                      System
                    </label>
                  </div>
                </div>
              </div>
            </div>
            
            <div>
              <h3 className="text-lg font-medium text-gray-900">Font Size</h3>
              <p className="mt-1 text-sm text-gray-500">
                Adjust the font size for better readability.
              </p>
              <div className="mt-4">
                <div className="w-full">
                  <input
                    type="range"
                    min="1"
                    max="5"
                    defaultValue="3"
                    className="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer"
                  />
                  <div className="flex justify-between text-xs text-gray-500 mt-1">
                    <span>Small</span>
                    <span>Medium</span>
                    <span>Large</span>
                  </div>
                </div>
              </div>
            </div>
            
            <div className="flex justify-end">
              <button
                type="button"
                className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              >
                Save Appearance
              </button>
            </div>
          </div>
        )}
        
        {activeTab === 'advanced' && (
          <div className="space-y-6">
            <div>
              <h3 className="text-lg font-medium text-gray-900">Advanced Settings</h3>
              <p className="mt-1 text-sm text-gray-500">
                Configure advanced settings for your Local AI Agent System.
              </p>
              <div className="mt-4 space-y-4">
                <div>
                  <label htmlFor="max-tokens" className="block text-sm font-medium text-gray-700">
                    Max Tokens
                  </label>
                  <div className="mt-1">
                    <input
                      type="number"
                      name="max-tokens"
                      id="max-tokens"
                      className="shadow-sm focus:ring-blue-500 focus:border-blue-500 block w-full sm:text-sm border-gray-300 rounded-md"
                      defaultValue="2048"
                    />
                  </div>
                  <p className="mt-1 text-sm text-gray-500">
                    Maximum number of tokens to generate in AI responses.
                  </p>
                </div>
                
                <div>
                  <label htmlFor="temperature" className="block text-sm font-medium text-gray-700">
                    Temperature
                  </label>
                  <div className="mt-1">
                    <input
                      type="number"
                      name="temperature"
                      id="temperature"
                      className="shadow-sm focus:ring-blue-500 focus:border-blue-500 block w-full sm:text-sm border-gray-300 rounded-md"
                      defaultValue="0.7"
                      step="0.1"
                      min="0"
                      max="2"
                    />
                  </div>
                  <p className="mt-1 text-sm text-gray-500">
                    Controls randomness in AI responses. Lower values are more deterministic.
                  </p>
                </div>
                
                <div>
                  <label htmlFor="model" className="block text-sm font-medium text-gray-700">
                    Default AI Model
                  </label>
                  <div className="mt-1">
                    <select
                      id="model"
                      name="model"
                      className="shadow-sm focus:ring-blue-500 focus:border-blue-500 block w-full sm:text-sm border-gray-300 rounded-md"
                    >
                      <option value="gpt-4">GPT-4</option>
                      <option value="gpt-3.5-turbo" selected>GPT-3.5 Turbo</option>
                      <option value="claude-2">Claude 2</option>
                      <option value="llama-2">Llama 2</option>
                    </select>
                  </div>
                  <p className="mt-1 text-sm text-gray-500">
                    Default AI model to use for tasks.
                  </p>
                </div>
              </div>
            </div>
            
            <div>
              <h3 className="text-lg font-medium text-gray-900 text-red-600">Danger Zone</h3>
              <p className="mt-1 text-sm text-gray-500">
                These actions are destructive and cannot be undone.
              </p>
              <div className="mt-4 space-y-4">
                <div className="flex justify-between items-center p-4 bg-red-50 rounded-md">
                  <div>
                    <h4 className="text-sm font-medium text-red-800">Clear All Data</h4>
                    <p className="text-sm text-red-700">
                      This will permanently delete all your tasks, files, and settings.
                    </p>
                  </div>
                  <button
                    type="button"
                    className="inline-flex items-center px-3 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
                  >
                    Clear Data
                  </button>
                </div>
                
                <div className="flex justify-between items-center p-4 bg-red-50 rounded-md">
                  <div>
                    <h4 className="text-sm font-medium text-red-800">Reset to Default</h4>
                    <p className="text-sm text-red-700">
                      This will reset all settings to their default values.
                    </p>
                  </div>
                  <button
                    type="button"
                    className="inline-flex items-center px-3 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
                  >
                    Reset
                  </button>
                </div>
              </div>
            </div>
            
            <div className="flex justify-end">
              <button
                type="button"
                className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              >
                Save Advanced Settings
              </button>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

export default Settings;
