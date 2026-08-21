import React from 'react';
import { STATUS_META } from '../constants.js';

export default function StatusChip({ status }) {
  const meta = STATUS_META[status] ?? { label: status, chipClass: 'status-chip' };
  return <span className={meta.chipClass}>{meta.label}</span>;
}
