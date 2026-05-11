import { useState } from "react";
import api from "../api/axios";

export default function Login() {
  const [mode, setMode] = useState("login"); 
  // login | register

  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  // ===== LOGIN =====
  const handleLogin = async () => {
    try {
      const res = await api.post("/auth/login", {
        username,
        password,
      });

      localStorage.setItem("token", res.data.token);
      alert("Login success");
    } catch (err) {
      alert("Login failed");
    }
  };

  // ===== REGISTER =====
  const handleRegister = async () => {
    try {
      const res = await api.post("/auth/register", {
        username,
        email,
        password,
      });

      alert("Register success → now login");
      setMode("login");
    } catch (err) {
      alert("Register failed");
    }
  };

  return (
    <div style={{ padding: 40 }}>
      <h2>{mode === "login" ? "Login" : "Register"}</h2>

      <input
        placeholder="username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />

      <br />

      <input
        placeholder="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />

      <br />

      <input
        placeholder="password"
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />

      <br />

      {/* ACTION BUTTON */}
      {mode === "login" ? (
        <button onClick={handleLogin}>Login</button>
      ) : (
        <button onClick={handleRegister}>Register</button>
      )}

      <br /><br />

      {/* SWITCH MODE */}
      {mode === "login" ? (
        <p>
          No account?{" "}
          <button onClick={() => setMode("register")}>
            Register here
          </button>
        </p>
      ) : (
        <p>
          Already have account?{" "}
          <button onClick={() => setMode("login")}>
            Login here
          </button>
        </p>
      )}
    </div>
  );
}