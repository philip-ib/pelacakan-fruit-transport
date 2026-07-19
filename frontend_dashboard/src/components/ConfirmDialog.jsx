import React from 'react';

export default function ConfirmDialog({ open, title, message, onConfirm, onCancel }) {
  if (!open) return null;

  return (
    <div style={{
      position: 'fixed', top: 0, left: 0, right: 0, bottom: 0,
      backgroundColor: 'rgba(0,0,0,0.4)', display: 'flex',
      alignItems: 'center', justifyContent: 'center', zIndex: 1000,
    }}>
      <div style={{
        backgroundColor: '#fff', borderRadius: '12px', padding: '24px',
        minWidth: '320px', maxWidth: '420px', boxShadow: '0 8px 32px rgba(0,0,0,0.2)',
      }}>
        <h3 style={{ margin: '0 0 8px', fontSize: '18px' }}>{title}</h3>
        <p style={{ margin: '0 0 20px', color: '#666', fontSize: '14px' }}>{message}</p>
        <div style={{ display: 'flex', justifyContent: 'flex-end', gap: '8px' }}>
          <button onClick={onCancel} style={{
            padding: '8px 16px', border: '1px solid #ccc',
            borderRadius: '6px', background: '#fff', cursor: 'pointer',
          }}>Batal</button>
          <button onClick={onConfirm} style={{
            padding: '8px 16px', border: 'none',
            borderRadius: '6px', background: '#d32f2f', color: '#fff', cursor: 'pointer',
          }}>Hapus</button>
        </div>
      </div>
    </div>
  );
}
