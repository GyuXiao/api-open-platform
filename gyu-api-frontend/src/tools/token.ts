export function setToken(token: string) {
  localStorage.setItem('jwt', token);
}

export function getToken(): string {
  return localStorage.getItem('jwt') || '';
}
