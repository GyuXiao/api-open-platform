export function setToken(token: string) {
  localStorage.setItem('token', token);
}

export function getToken(): string {
  return localStorage.getItem('token') || '';
}
