const BASE_URL = '/api/v1';

async function request(path, { method = 'GET', body } = {}) {
  try {
    const res = await fetch(`${BASE_URL}${path}`, {
      method,
      headers: body ? { 'Content-Type': 'application/json' } : undefined,
      body: body ? JSON.stringify(body) : undefined,
    });
    const json = await res.json();
    return { success: json.success === true, message: json.message ?? null, data: json.data ?? null };
  } catch (e) {
    return { success: false, message: `Gagal terhubung ke server: ${e.message}` };
  }
}

export const getAllTransports = () => request('/transports');

export const updateStatus = (id, status) =>
  request(`/transports/${id}/status`, { method: 'PATCH', body: { status } });

export const deleteTransport = (id) =>
  request(`/transports/${id}`, { method: 'DELETE' });

export function connectWebSocket(onMessage) {
  const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
  const host = window.location.host;
  const ws = new WebSocket(`${protocol}//${host}/ws`);

  ws.onmessage = () => onMessage();
  ws.onerror = () => {}; // akan trigger onclose
  ws.onclose = () => onMessage('reconnect');

  return ws;
}
