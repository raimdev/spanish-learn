import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { AuthProvider } from './context/AuthContext';
import ProtectedRoute from './components/ProtectedRoute';
import Layout from './components/Layout';
import LoginPage from './pages/LoginPage';
import DashboardPage from './pages/DashboardPage';
import GrammarListPage from './pages/GrammarListPage';
import GrammarDetailPage from './pages/GrammarDetailPage';
import VocabularyPage from './pages/VocabularyPage';
import ExercisesListPage from './pages/ExercisesListPage';
import ExercisePage from './pages/ExercisePage';
import FlashcardsPage from './pages/FlashcardsPage';
import ProgressPage from './pages/ProgressPage';
import './styles/index.css';

const router = createBrowserRouter([
  {
    path: '/login',
    element: <LoginPage />,
  },
  {
    element: <ProtectedRoute />,
    children: [
      {
        element: <Layout />,
        children: [
          { path: '/', element: <DashboardPage /> },
          { path: '/grammar', element: <GrammarListPage /> },
          { path: '/grammar/:slug', element: <GrammarDetailPage /> },
          { path: '/vocabulary', element: <VocabularyPage /> },
          { path: '/exercises', element: <ExercisesListPage /> },
          { path: '/exercises/:id', element: <ExercisePage /> },
          { path: '/flashcards', element: <FlashcardsPage /> },
          { path: '/progress', element: <ProgressPage /> },
        ],
      },
    ],
  },
]);

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <AuthProvider>
      <RouterProvider router={router} />
    </AuthProvider>
  </React.StrictMode>
);
