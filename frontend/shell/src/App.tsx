import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { Suspense, lazy } from 'react';
import MainLayout from './layouts/MainLayout';
import LoadingFallback from './components/LoadingFallback';
import ErrorBoundary from './components/ErrorBoundary';

// Lazy load micro frontends
const TaskManagement = lazy(() => import('taskManagement/App'));
const ResultViewer = lazy(() => import('resultViewer/App'));

function App() {
  return (
    <Router>
      <MainLayout>
        <ErrorBoundary>
          <Suspense fallback={<LoadingFallback />}>
            <Routes>
              <Route path="/tasks" element={<TaskManagement />} />
              <Route path="/results" element={<ResultViewer />} />
              <Route path="/" element={<Navigate to="/tasks" replace />} />
            </Routes>
          </Suspense>
        </ErrorBoundary>
      </MainLayout>
    </Router>
  );
}

export default App;
