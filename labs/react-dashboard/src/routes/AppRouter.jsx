import { BrowserRouter, Routes, Route } from "react-router";

import PrivateRoute from "./PrivateRoute";
import ProtectedRoute from "./ProtectedRoute";
import Register from "../features/register/Register";
import Login from "../features/login/Login";
import Home from "../features/home/Home";

const AppRouter = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<PrivateRoute />}>
          <Route path="/" element={<Home />} />
        </Route>

        <Route path="/" element={<ProtectedRoute />}>
          <Route path="/register" element={<Register />} />
          <Route path="/login" element={<Login />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
};

export default AppRouter;
