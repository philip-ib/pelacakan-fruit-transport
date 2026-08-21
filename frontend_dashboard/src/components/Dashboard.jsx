import React, { useState, useEffect, useRef, useCallback } from 'react';
import { getAllTransports, updateStatus, deleteTransport, connectWebSocket } from '../services/api.js';
import { STATUS_DALAM_PERJALANAN, STATUS_DITERIMA } from '../constants.js';
import StatusChip from './StatusChip.jsx';
import ConfirmDialog from './ConfirmDialog.jsx';

function formatDate(iso) {
  try {
    const d = new Date(iso);
    return `${d.getDate()}/${d.getMonth() + 1}/${d.getFullYear()} ${String(d.getHours()).padStart(2, '0')}:${String(d.getMinutes()).padStart(2, '0')}`;
  } catch { return iso; }
}

export default function Dashboard() {
  const [transports, setTransports] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [updatingId, setUpdatingId] = useState(null);
  const [deletingId, setDeletingId] = useState(null);
  const [confirmDelete, setConfirmDelete] = useState(null);
  const wsRef = useRef(null);

  const fetchData = useCallback(async () => {
    const result = await getAllTransports();
    if (result.success) {
      setTransports(result.data);
      setError(null);
    } else {
      setError(result.message ?? 'Gagal mengambil data');
    }
    setLoading(false);
  }, []);

  useEffect(() => {
    fetchData();
  }, [fetchData]);

  // WebSocket real-time
  useEffect(() => {
    let reconnectTimer;
    const connect = () => {
      wsRef.current = connectWebSocket((type) => {
        if (type === 'reconnect') {
          reconnectTimer = setTimeout(connect, 3000);
        } else {
          fetchData();
        }
      });
    };
    connect();
    return () => {
      clearTimeout(reconnectTimer);
      wsRef.current?.close();
    };
  }, [fetchData]);

  const runMutation = async (setPendingId, id, mutation, onSuccess) => {
    setPendingId(id);
    const res = await mutation();
    setPendingId(null);
    if (res.success) {
      onSuccess();
    } else {
      alert(res.message);
    }
  };

  const handleTerima = (t) =>
    runMutation(setUpdatingId, t.id, () => updateStatus(t.id, STATUS_DITERIMA), fetchData);

  const handleDeleteRequest = (t) => {
    setConfirmDelete(t);
  };

  const handleDeleteConfirm = () => {
    const t = confirmDelete;
    setConfirmDelete(null);
    runMutation(setDeletingId, t.id, () => deleteTransport(t.id), fetchData);
  };

  // --- Loading ---
  if (loading && transports.length === 0) {
    return (
      <div className="state-box">
        <div className="spinner" />
        <p>Memuat data...</p>
      </div>
    );
  }

  // --- Error ---
  if (error && transports.length === 0) {
    return (
      <div className="state-box">
        <div className="icon-big">☁️</div>
        <p className="error-text">{error}</p>
        <button onClick={fetchData} className="btn btn-green">🔄 Coba Lagi</button>
      </div>
    );
  }

  // --- Empty ---
  if (transports.length === 0) {
    return (
      <div className="state-box">
        <div className="icon-big">📥</div>
        <p>Belum ada data pengangkutan</p>
      </div>
    );
  }

  return (
    <>
      <div className="table-wrapper">
        <table>
          <thead>
            <tr>
              <th>No</th>
              <th>Nomor Truk</th>
              <th>ID TPH</th>
              <th>Berat (Ton)</th>
              <th>Status</th>
              <th>Waktu Dibuat</th>
              <th>Aksi</th>
            </tr>
          </thead>
          <tbody>
            {transports.map((t, i) => {
              const isDalamPerjalanan = t.status === STATUS_DALAM_PERJALANAN;
              return (
                <tr key={t.id}>
                  <td>{i + 1}</td>
                  <td>{t.nomor_truk}</td>
                  <td>{t.id_tph}</td>
                  <td>{Number(t.berat_estimasi).toFixed(1)}</td>
                  <td><StatusChip status={t.status} /></td>
                  <td>{formatDate(t.created_at)}</td>
                  <td className="action-cell">
                    {isDalamPerjalanan && (
                      <button
                        className="btn btn-green btn-sm"
                        disabled={updatingId === t.id}
                        onClick={() => handleTerima(t)}
                      >
                        {updatingId === t.id ? '...' : '✅ Terima Buah'}
                      </button>
                    )}
                    <button
                      className="btn btn-danger btn-icon"
                      disabled={deletingId === t.id || updatingId === t.id}
                      onClick={() => handleDeleteRequest(t)}
                      title="Hapus data"
                    >
                      {deletingId === t.id ? '⟳' : '🗑️'}
                    </button>
                  </td>
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>

      <ConfirmDialog
        open={!!confirmDelete}
        title="Konfirmasi Hapus"
        message={`Hapus data pengangkutan dari truk ${confirmDelete?.nomor_truk}?`}
        onConfirm={handleDeleteConfirm}
        onCancel={() => setConfirmDelete(null)}
      />
    </>
  );
}
