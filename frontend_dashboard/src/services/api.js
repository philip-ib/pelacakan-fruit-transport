const BASE_URL = '/api/v1';

export async function getAllTransports() {
  try {
    const res = await fetch(`${BASE_URL}/transports`);
    const json = await res.json();
    if (json.success && json.data) {
      return { success: true, data: json.data };
    }
    return { success: false, message: json.message || 'Gagal mengambil data' };
  } catch (e) {
    return { success: false, message: `Gagal terhubung ke server: ${e.message}` };
  }
}

export async function updateStatus(id, status) {
  try {
    const res = await fetch(`${BASE_URL}/transports/${id}/status`, {
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ status }),
    });
    const json = await res.json();
    return { success: json.success, message: json.message };
  } catch (e) {
    return { success: false, message: `Gagal terhubung ke server: ${e.message}` };
  }
}

export async function deleteTransport(id) {
  try {
    const res = await fetch(`${BASE_URL}/transports/${id}`, {
      method: 'DELETE',
      headers: { 'Content-Type': 'application/json' },
    });
    const json = await res.json();
    return { success: json.success, message: json.message };
  } catch (e) {
    return { success: false, message: `Gagal terhubung ke server: ${e.message}` };
  }
}

export function connectWebSocket(onMessage) {
  const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
  const host = window.location.host;
  const ws = new WebSocket(`${protocol}//${host}/ws`);

  ws.onmessage = () => onMessage();
  ws.onerror = () => {}; // akan trigger onclose
  ws.onclose = () => onMessage('reconnect');

  return ws;
}
