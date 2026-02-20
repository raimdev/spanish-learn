import { NavLink, Outlet } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import { VocabularyProvider } from '../context/VocabularyContext';

export default function Layout() {
  const { user, logout } = useAuth();

  return (
    <VocabularyProvider>
    <div className="app">
      <nav className="navbar">
        <div className="nav-brand">
          <NavLink to="/">Spanish Learn</NavLink>
        </div>
        <div className="nav-links">
          <NavLink to="/" end>Dashboard</NavLink>
          <NavLink to="/grammar">Grammar</NavLink>
          <NavLink to="/vocabulary">Vocabulary</NavLink>
          <NavLink to="/exercises">Exercises</NavLink>
          <NavLink to="/flashcards">Flashcards</NavLink>
          <NavLink to="/progress">Progress</NavLink>
        </div>
        <div className="nav-user">
          <span>{user?.username}</span>
          <button onClick={logout} className="btn-logout">Logout</button>
        </div>
      </nav>
      <main className="main-content">
        <Outlet />
      </main>
    </div>
    </VocabularyProvider>
  );
}
