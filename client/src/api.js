let token = null;

const setToken = _token => {
  token = _token;
};

const headers = () => {
  const _headers = {
    Accept: "application/json",
    "Content-Type": "application/json"
  };
  if (token) {
    _headers["Authorization"] = `Bearer ${token}`;
  }
  return _headers;
};

const currentUser = async () => {
  const response = await fetch("/api/user", {
    headers: headers()
  });
  return response.json();
};

const signup = async (username, password) => {
  const response = await fetch("/api/signup", {
    method: "POST",
    headers: headers(),
    body: JSON.stringify({
      username,
      password
    })
  });
  return response.json();
};

const login = async (username, password) => {
  const response = await fetch("/api/login", {
    method: "POST",
    headers: headers(),
    body: JSON.stringify({
      username,
      password
    })
  });
  return response.json();
};

const newProgram = async program => {
  const response = await fetch("/api/program", {
    method: "POST",
    headers: headers(),
    body: JSON.stringify(program)
  });
  return response.json();
};

export default { setToken, currentUser, login, signup, newProgram };
