import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { ToastContainer } from "react-toastify";

import "./index.css";
import AppRouter from "./routes/AppRouter.jsx";

createRoot(document.getElementById("root")).render(
  <StrictMode>
    <AppRouter />
    <ToastContainer
      position="top-right"
      autoClose={3000}
      pauseOnHover
      theme="light"
    />
  </StrictMode>
);
