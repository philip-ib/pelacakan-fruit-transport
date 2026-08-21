import React from 'react';

export default function ConfirmDialog({
  open, title, message, confirmLabel = 'Hapus', onConfirm, onCancel,
}) {
  if (!open) return null;

  return (
    <div className="dialog-overlay">
      <div className="dialog">
        <h3 className="dialog-title">{title}</h3>
        <p className="dialog-message">{message}</p>
        <div className="dialog-actions">
          <button onClick={onCancel} className="btn btn-outline">Batal</button>
          <button onClick={onConfirm} className="btn btn-danger-solid">{confirmLabel}</button>
        </div>
      </div>
    </div>
  );
}
