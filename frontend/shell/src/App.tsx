import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import MainLayout from './layouts/MainLayout';
import ErrorBoundary from './components/ErrorBoundary';
import Dashboard from './components/Dashboard';
import CodeExecution from './components/CodeExecution';
import WebBrowsing from './components/WebBrowsing';
import Filesystem from './components/Filesystem';
import Settings from './components/Settings';

// Micro frontend components using iframes
const TaskManagement = () => (
  <div className="h-full">
    <iframe
      src="http://localhost:5001"
      title="Task Management"
      className="w-full h-full border-none"
      style={{ minHeight: '600px' }}
    />
  </div>
);

const ResultViewer = () => (
  <div className="h-full">
    <iframe
      src="http://localhost:5002"
      title="Result Viewer"
      className="w-full h-full border-none"
      style={{ minHeight: '600px' }}
    />
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
