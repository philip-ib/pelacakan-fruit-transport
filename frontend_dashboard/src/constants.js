export const STATUS_DALAM_PERJALANAN = 'DALAM_PERJALANAN';
export const STATUS_DITERIMA = 'DITERIMA';

// Label tampilan + kelas chip untuk setiap status (satu sumber kebenaran).
export const STATUS_META = {
  [STATUS_DALAM_PERJALANAN]: { label: 'Dalam Perjalanan', chipClass: 'status-chip' },
  [STATUS_DITERIMA]: { label: 'Diterima', chipClass: 'status-chip status-chip--diterima' },
};
