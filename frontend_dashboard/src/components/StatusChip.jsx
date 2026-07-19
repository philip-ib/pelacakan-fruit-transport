import React from 'react';

export default function StatusChip({ status }) {
  const isDiterima = status === 'DITERIMA';
  return (
    <span style={{
      display: 'inline-block',
      padding: '4px 12px',
      borderRadius: '16px',
      fontSize: '12px',
      fontWeight: 600,
      color: isDiterima ? '#2e7d32' : '#e65100',
      backgroundColor: isDiterima ? '#e8f5e9' : '#fff3e0',
      border: `1px solid ${isDiterima ? '#a5d6a7' : '#ffcc02'}`,
    }}>
      {status}
    </span>
  );
}
