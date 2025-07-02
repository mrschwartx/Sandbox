import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { ThemeProvider } from '@mui/material';
import { ToastContainer } from "react-toastify";

import "./index.css";
import theme from "./themes/theme.jsx";
import AppRouter from "./routes/AppRouter.jsx";

createRoot(document.getElementById("root")).render(
  <StrictMode>
    <ThemeProvider theme={theme}>
      <AppRouter />
      <ToastContainer
        position="top-right"
        autoClose={3000}
        pauseOnHover
        theme="light"
      />
    </ThemeProvider>
  </StrictMode>
);
