import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import MainLayout from './layouts/MainLayout';
import ErrorBoundary from './components/ErrorBoundary';
import Dashboard from './components/Dashboard';
import CodeExecution from './components/CodeExecution';
import WebBrowsing from './components/WebBrowsing';
import Filesystem from './components/Filesystem';
import Settings from './components/Settings';

// Placeholder components for micro frontends
const TaskManagement = () => (
  <div className="p-8 bg-white shadow-md rounded-lg">
    <h2 className="text-2xl font-bold mb-4">Task Management</h2>
    <p className="text-gray-600">
      This is a placeholder for the Task Management application.
      In a production environment, this would be loaded as a micro frontend.
    </p>
    <div className="mt-4 p-4 bg-blue-50 border border-blue-200 rounded-md">
      <p className="text-blue-700">
        The micro frontend integration is currently being configured.
        Please check back later.
      </p>
    </div>
  </div>
);

const ResultViewer = () => (
  <div className="p-8 bg-white shadow-md rounded-lg">
    <h2 className="text-2xl font-bold mb-4">Result Viewer</h2>
    <p className="text-gray-600">
      This is a placeholder for the Result Viewer application.
      In a production environment, this would be loaded as a micro frontend.
    </p>
    <div className="mt-4 p-4 bg-blue-50 border border-blue-200 rounded-md">
      <p className="text-blue-700">
        The micro frontend integration is currently being configured.
        Please check back later.
      </p>
    </div>
  </div>
);

function App() {
  return (
    <Router>
      <MainLayout>
        <ErrorBoundary>
          <Routes>
            <Route path="/" element={<Dashboard />} />
            <Route path="/tasks" element={<TaskManagement />} />
            <Route path="/results" element={<ResultViewer />} />
            <Route path="/code-execution" element={<CodeExecution />} />
            <Route path="/web-browsing" element={<WebBrowsing />} />
            <Route path="/filesystem" element={<Filesystem />} />
            <Route path="/settings" element={<Settings />} />
            <Route path="*" element={<Navigate to="/" replace />} />
          </Routes>
        </ErrorBoundary>
      </MainLayout>
    </Router>
  );
}

export default App;
