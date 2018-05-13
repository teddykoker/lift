let token = null;

const setToken = _token => {
  token = _token;
};

const headers = () => {
  const _headers = {
    'Accept': 'application/json',
    'Content-Type': 'application/json'
  }
  if (token) { _headers["Authorization"] = `Bearer ${token}` }
}


const currentUser = async () => {
  const response = await fetch('/api/user', {
    headers: headers()
  })
  return response.json()
}

const login = async (username, password) => {
  const response = await fetch('/api/login', {
    method: "POST",
    headers: headers(),
    body: JSON.stringify({
      username,
      password
    })
  })
  return response.json()
}

export default { setToken, currentUser, login };
