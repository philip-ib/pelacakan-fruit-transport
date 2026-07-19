import React from 'react';
import Dashboard from './components/Dashboard.jsx';

export default function App() {
  return (
    <div className="app">
      <header className="app-header">
        <h1>🚛 Dashboard Admin PKS</h1>
        <span className="subtitle">Pelacakan Pengangkutan Buah Sawit TPH → PKS</span>
      </header>
      <main className="app-main">
        <Dashboard />
      </main>
    </div>
  );
}
